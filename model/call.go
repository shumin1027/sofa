package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"os/user"
	"strings"
	"time"
	"xtc/sofa/connect"
	"xtc/sofa/log"
	"xtc/sofa/pkg/shell"
)

/*
一次 shell 调用的 结构封装
*/
type Call struct {
	// 请求id
	TID string `json:"TID" gorm:"column:tid"`
	// 平台 LSF or SLURM
	Platform string `json:"Platform" gorm:"column:platform"`
	// 执行命令
	Command string `json:"Command" gorm:"column:command"`
	// 完整的执行命令
	FullCommand string `json:"FullCommand" gorm:"column:full_command"`
	// 退出状态
	ExitStatus int `json:"ExitStatus" gorm:"column:exit_status"`
	// 命令执行结果：标准输出
	Stdout []string `json:"Stdout" gorm:"column:stdout"`
	// 命令执行结果：标准错误输出
	Stderr []string `json:"Stderr" gorm:"column:stderr"`
	// 执行命令的linux用户
	Username string `json:"Username" gorm:"column:username"`
	// 任务提交时间
	SubmitTime time.Time `json:"SubmitTime" gorm:"column:submit_time"`
	// 执行开始时间
	StartTime time.Time `json:"StartTime" gorm:"column:start_time"`
	// 执行完成时间
	EndTime time.Time `json:"EndTime" gorm:"column:end_time"`
}

/*
执行命令
*/
func (call *Call) Exec() {

	cmd := call.FullCommand

	// 当前linux用户
	u, err := user.Current()

	if err != nil {
		log.Logger.Error("get current user failed", zap.Error(err))
	}

	if len(call.Username) > 0 && call.Username != u.Username {
		cmd = fmt.Sprintf("su %s -c '%s'", call.Username, cmd)
	} else {
		call.Username = u.Username
	}

	call.StartTime = time.Now()
	pipe := shell.Exec(cmd)
	call.EndTime = time.Now()

	if pipe.Error() != nil {
		log.Logger.Error("exec error", zap.String("command", cmd), zap.Error(pipe.Error()))
	}

	exit := pipe.ExitStatus()
	call.ExitStatus = exit

	scanner := bufio.NewScanner(pipe.Reader)
	if exit != 0 {
		for scanner.Scan() {
			line := scanner.Text()
			call.Stderr = append(call.Stdout, line)
		}
	} else {
		for scanner.Scan() {
			line := scanner.Text()
			call.Stdout = append(call.Stdout, line)
		}
	}

}

/*
存储到Redis
*/
func (call *Call) Save() {
	// json 编码
	j, err := json.Marshal(call)
	if err != nil {
		log.Logger.Error("marshal json error", zap.Error(err))
	}

	// 存入redis给logstash继续处理
	redis := connect.RedisClient()
	redis.LPush(strings.ToLower(call.Platform+"-"+call.Command), j)
	log.Logger.Info("push to redis successed")

}

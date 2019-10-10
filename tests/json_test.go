package tests

import (
	"encoding/json"
	"fmt"
	"os/user"
	"testing"
	"time"
	"xtc/sofa/model"
)

// 一次 shell 调用的 结构封装
type Call struct {
	// 请求id
	TID string `json:"TID" gorm:"column:tid"`
	// 平台 LSF or SLURM
	Platform string `json:"Platform" gorm:"column:platform"`
	// 执行命令
	Command string `json:"Command" gorm:"column:command"`
	// 退出状态
	ExitStatus int `json:"ExitStatus" gorm:"column:exit_status"`
	// 命令执行结果：标准输出
	Stdout []string `json:"Stdout" gorm:"column:stdout"`
	// 命令执行结果：标准错误输出
	Stderr []string `json:"Stderr" gorm:"column:stderr"`
	// 执行命令的linux用户
	User string `json:"User" gorm:"column:user"`
	// 执行完成时间
	Time time.Time `json:"Time" gorm:"column:time"`
}

func TestMarshal(t *testing.T) {

	u, _ := user.Current()
	call := new(model.Call)
	call.TID = "001"
	call.Platform = "docker"
	call.Command = "docker ps"
	call.Time = time.Now()
	call.Stdout = make([]string, 0, 10)
	call.User = u.Username

	line := `"docker-entrypoint.s…"`

	call.Stdout = append(call.Stdout, line)

	js, _ := json.Marshal(call)

	fmt.Println(string(js))

}

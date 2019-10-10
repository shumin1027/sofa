package model

import "time"

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
	Username string `json:"Username" gorm:"column:username"`
	// 执行完成时间
	Time time.Time `json:"Time" gorm:"column:time"`
}

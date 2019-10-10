package cmd

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
	. "xtc/sofa/log"
	"xtc/sofa/model"
	"xtc/sofa/pkg/socket/client"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec COMMAND",
	Short: "执行一条Shell命令，将命令输出结果作为输入",
	Long:  `执行一条Shell命令，将命令输出结果作为输入`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// 初始化日志配置
		InitLogger("agent")
	},
	Run: func(cmd *cobra.Command, args []string) {

		// 初始化参数
		p := new(parameter)
		tid, err := cmd.Flags().GetString("tid")
		if err != nil {
		}
		p.tid = tid

		platform, err := cmd.Flags().GetString("platform")
		if err != nil {
		}
		p.platform = platform

		command, err := cmd.Flags().GetString("command")
		if err != nil {
		}
		p.command = command

		username, err := cmd.Flags().GetString("username")
		if err != nil {
		}
		p.username = username

		if len(args) != 1 {
			fmt.Println("exec needs one command")
		}

		exec(p, args[0])
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.Flags().StringP("tid", "t", "", "TraceId,用于跟踪任务执行")
	execCmd.Flags().StringP("platform", "p", "SLURM", "作业调度平台：LSF/SLURM")
	execCmd.Flags().StringP("command", "c", "", "执行的命令,如：bjobs")
	execCmd.Flags().StringP("username", "u", "", "运行命令使用的Linux用户")
}

// 执行 cmd 命令 作为输入
func exec(p *parameter, cmd string) error {

	Logger.Info("exec:" + cmd)

	call := new(model.Call)
	call.TID = p.tid
	call.Platform = p.platform
	call.Command = p.command
	call.FullCommand = cmd
	call.SubmitTime = time.Now()
	call.Stdout = make([]string, 0, 10)

	if len(p.username) > 0 {
		call.Username = p.username
	}

	call.Exec()

	// 序列化
	var datas bytes.Buffer
	encoder := gob.NewEncoder(&datas)
	encoder.Encode(call)

	// 使用 unix socket 发送给 server 处理
	client.SentWithBytes(datas.Bytes())

	return nil
}

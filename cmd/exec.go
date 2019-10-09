/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/bitfield/script"
	"os/user"
	"time"
	"xtc/sofa/model"
	"xtc/sofa/pkg/socket/client"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec COMMAND",
	Short: "执行一条Shell命令，将命令输出结果作为输入",
	Long:  `执行一条Shell命令，将命令输出结果作为输入`,
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

		user, err := cmd.Flags().GetString("user")
		if err != nil {
		}
		p.user = user

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
	execCmd.Flags().StringP("user", "u", "", "运行命令使用的Linux账户")
}

// 执行 cmd 命令 作为输入
func exec(p *parameter, cmd string) {

	// 当前linux用户
	u, _ := user.Current()

	call := new(model.Call)
	call.TID = p.tid
	call.Platform = p.platform
	call.Command = p.command
	call.Time = time.Now()
	call.Stdout = make([]string, 0, 10)
	call.User = u.Username

	if p.user != "" {
		cmd = fmt.Sprintf("su %s -c '%s' ", p.user, cmd)
		call.User = p.user
	}

	pipe := script.Exec(cmd)

	exit := pipe.ExitStatus()
	if exit != 0 {

	}
	call.ExitStatus = exit

	scanner := bufio.NewScanner(pipe.Reader)

	for scanner.Scan() {
		line := scanner.Text()
		call.Stdout = append(call.Stdout, line)
	}

	// 序列化
	var datas bytes.Buffer
	encoder := gob.NewEncoder(&datas)
	encoder.Encode(call)

	// 使用 unix socket 发送给 server 处理
	client.SentWithBytes(datas.Bytes())

}

/*
Copyright © 2019 shumin

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
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"time"
	"xtc/sofa/log"
	"xtc/sofa/model"
	"xtc/sofa/pkg/socket/client"
	"xtc/sofa/pkg/version"
)

var cfgFile string

// 执行参数
type parameter struct {
	tid      string // TraceId,用于跟踪任务执行
	platform string // 作业调度平台：LSF/SLURM
	command  string // 执行的命令,如：bjobs
	username string // 执行的命令的用户,如：root
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sofa",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PreRun: func(cmd *cobra.Command, args []string) {
		// 初始化日志配置
		log.InitLogger("agent")
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

		run(p)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.AddCommand(version.Command())
	rootCmd.Flags().StringP("tid", "t", "", "TraceId,用于跟踪任务执行")
	rootCmd.Flags().StringP("platform", "p", "SLURM", "作业调度平台：LSF/SLURM")
	rootCmd.Flags().StringP("command", "c", "", "执行的命令,如：bjobs")
}

// 从 管道 接受输入
func run(p *parameter) {

	// 当前linux用户
	u, _ := user.Current()

	call := new(model.Call)
	call.TID = p.tid
	call.Platform = p.platform
	call.Command = p.command
	call.Username = u.Username
	call.EndTime = time.Now()
	call.Stdout = make([]string, 0, 10)

	info, _ := os.Stdin.Stat()

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage:")
		fmt.Println("bjobs -W | sofa --platform=LSF --command=bjobs --tid=001")
		fmt.Println("or")
		fmt.Println("sofa exec 'bjobs -W' --platform=LSF --command=bjobs --tid=001")
		os.Exit(1)
	}

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		line := s.Text()
		call.Stdout = append(call.Stdout, line)
	}

	// 序列化
	var datas bytes.Buffer
	encoder := gob.NewEncoder(&datas)
	encoder.Encode(call)

	// 使用 unix socker 发送给 server 处理
	client.SentWithBytes(datas.Bytes())

}

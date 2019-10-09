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
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"xtc/sofa/connect"
	"xtc/sofa/pkg/socket/server"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动一个server进程，通过Unix Socket来获取sofa收集到到数据,然后转发到Redis",
	Long:  `启动一个server进程，通过Unix Socket来获取sofa收集到到数据,然后转发到Redis`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// init redis
		address := viper.GetStringSlice("redis.address")
		db := viper.GetInt("redis.db")
		rconfg := &connect.RedisConfig{
			Address: address,
			DB:      db,
		}
		connect.InitRedis(rconfg)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server is starting……")
		server.Start()
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(serverCmd)
	serverCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.sofa.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".sofa" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".sofa")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
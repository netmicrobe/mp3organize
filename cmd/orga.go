// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	
	"github.com/spf13/cobra"
	"github.com/netmicrobe/mp3organize/utils"
)

// orgaCmd represents the orga command
var orgaCmd = &cobra.Command{
	Use:   "orga",
	Short: "根据id3v2信息重命名mp3文件：歌名-演唱人.mp3",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("错误！需要指明文件 或 文件夹的路径！")
			os.Exit(1)
		}
		orga(args[0])
	},
}

func orga(target string) {
	if target == "" {
		fmt.Println("错误！路径不能为空！")
		os.Exit(1)
	}
	// 重命名mp3文件
	utils.EachFiles(target, utils.Renamemp3)
}

func init() {
	rootCmd.AddCommand(orgaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// orgaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// orgaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

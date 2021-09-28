/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"care-screenshot/public"
	"io/ioutil"
	"log"
	"runtime"
	"strconv"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "给我一个URL，我截图成功发给企业微信机器人🤖",
	Long:  `命令行工具，可使用此工具订阅一些你关心的网页服务状态，然后添加到定时任务中。`,
	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flags().Lookup("url").Value.String()
		element := cmd.Flags().Lookup("element").Value.String()
		width, _ := strconv.Atoi(cmd.Flags().Lookup("kuan").Value.String())
		height, _ := strconv.Atoi(cmd.Flags().Lookup("gao").Value.String())
		bot := cmd.Flags().Lookup("bot").Value.String()

		launch := launcher.New().Headless(true)
		if runtime.GOOS == "darwin" {
			launch = launch.Bin(`/Applications/Google Chrome.app/Contents/MacOS/Google Chrome`)
		}
		if runtime.GOOS == "linux" {
			launch = launch.Set("--no-sandbox")
		}
		page := rod.New().ControlURL(launch.MustLaunch()).MustConnect().MustPage()
		defer page.Close()
		withTimeout := page.Timeout(time.Second)

		withTimeout.
			MustSetViewport(width, height, 1, false).
			MustNavigate(url).
			MustWaitLoad().MustWindowMaximize()
		// 这句很关键,能够控制程序等待页面渲染完毕之后再截图
		page.WaitRequestIdle(time.Duration(time.Second*10), []string{}, []string{})()
		el := page.MustElement(element).MustScreenshot()
		if err := ioutil.WriteFile("tmp.png", el, 0o644); err != nil {
			log.Fatal(err)
		}
		public.SendImage("tmp.png", bot)
	},
}

func init() {
	fset := execCmd.Flags()
	fset.StringP("url", "u", "https://baidu.com", "给我一个你想要截图的URL")
	fset.StringP("element", "e", "#s_lg_img", "给我你关心的页面元素")
	fset.StringP("kuan", "k", "1200", "页面宽度")
	fset.StringP("gao", "g", "800", "页面高度")
	fset.StringP("bot", "b", "d63e3f22-3a88-43fb-a2ad-ad78ba5b43b5", "机器人地址")
	execCmd.MarkFlagRequired("url")
	execCmd.MarkFlagRequired("element")
	rootCmd.AddCommand(execCmd)
}

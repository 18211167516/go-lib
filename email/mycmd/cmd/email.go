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
	"fmt"
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	To  []string
	Cc  []string
	Bcc []string
	//主题
	Subject string
	//内容
	Text string
	//html内容（优先）
	HTML string
	//抄送附件
	AttachFileName string
	From           string
	Smtp           string
	Port           int
	Password       string
)

// emailCmd represents the email command
var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "发送邮件",
	Long:  `为了更便捷的发送邮件`,
	Example: `
	基础发送：mycmd email -c "内容" -t "baichonghua@urthink.com"
	抄送：mycmd email -c "签到" -t "baichonghua@urthink.com" --cc "baichonghua@tansun.com.cn"
	私密抄送：mycmd email -c "签到" -t "baichonghua@urthink.com" --bcc "baichonghua@tansun.com.cn"
	发送html：mycmd email -c "签到" -t "baichonghua@urthink.com" --html='<a href="http://www.baidu.com">点击</a>'
	发送附件：mycmd email -c "签到" -t "baichonghua@urthink.com" --file="./cmd.toml.example"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		send()
	},
}

func init() {
	emailCmd.Flags().StringSliceVarP(&To, "to", "t", []string{}, "send email to")
	emailCmd.Flags().StringSliceVar(&Cc, "cc", []string{}, "send email Cc")
	emailCmd.Flags().StringSliceVar(&Bcc, "bcc", []string{}, "send email Bcc")
	emailCmd.Flags().StringVarP(&Subject, "subject", "s", "默认主题", "send email Subject")
	emailCmd.Flags().StringVarP(&Text, "context", "c", "", "send email Text")
	emailCmd.Flags().StringVar(&HTML, "html", "", "send email HTML contect")
	emailCmd.Flags().StringVar(&AttachFileName, "file", "", "send email AttachFile")
	emailCmd.MarkFlagRequired("to")
	emailCmd.MarkFlagRequired("context")
	rootCmd.AddCommand(emailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// emailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// emailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func send() {
	From = viper.GetString("email.from")
	Smtp = viper.GetString("email.smtp")
	Port = viper.GetInt("email.port")
	Password = viper.GetString("email.password")

	e := email.NewEmail()
	e.From = From
	e.To = To
	e.Cc = Cc
	e.Bcc = Bcc
	e.Subject = Subject
	e.AttachFile(AttachFileName)
	if HTML != "" {
		e.HTML = []byte(HTML)
	} else {
		e.Text = []byte(Text)
	}
	auth := smtp.PlainAuth("", From, Password, Smtp)
	addr := fmt.Sprintf("%s:%d", Smtp, Port)
	err := e.Send(addr, auth)
	if err != nil {
		log.Fatal("email ", err)
	}

	log.Println("发送成功")
}

package main

import (
	"chatroom/client/process"
	"fmt"
	"os"
)

//定义两个变量，一个表示用户id，一个表示用户密码
var userId int
var userPwd string

func main() {
	//接受用户的选择
	var key int
	//判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("---------欢迎登录多人聊天系统--------")
		fmt.Println("\t\t\t 1 登录聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择（1-3）：")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("USERID：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("PASSWORD：")
			fmt.Scanf("%s\n", &userPwd)
			//完成登录，创建一个UserProcess实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}
	}
}
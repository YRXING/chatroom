package process

import (
	"chatroom/client/model"
	"chatroom/common/message"
	"fmt"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User,10)
var CurUser model.CurUser //我们在用户登录成功后完成对CurUser的初始化

//在客户端显示当前在线的用户
func outputOnlineUser(){
	//如果我们要求不显示自己在线，下面我们增加一行代码
	//if v == userId{
	//	continue
	//}
	//遍历一把onlineUsers
	fmt.Println("当前在线用户列表：")
	for id,_:= range onlineUsers{
		fmt.Println("用户id：\t",id)
	}
}

//编写一个方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes){
	//适当优化
	user,ok:=onlineUsers[notifyUserStatusMes.UserId]
	if !ok{//如果map没有这个user，就新建一个
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.Status
	onlineUsers[notifyUserStatusMes.UserId] = user
	outputOnlineUser()
}


package process

import (
	"chatroom/client/utils"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的界面
func ShowMenu()  {
	fmt.Println("-------恭喜xxx登录成功------")
	fmt.Println("-------1. 显示在线用户列表------")
	fmt.Println("-------2. 发送消息------")
	fmt.Println("-------3. 信息列表------")
	fmt.Println("-------4. 退出系统------")
	fmt.Println("请选择（1-4）：")
	var key int
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
			fmt.Println("显示在线用户列表")
		case 2:
			fmt.Println("发送消息")
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你选择了退出系统...")
			os.Exit(0)
		default:
			fmt.Println("你输入的选项不正确...")
	}
}

//和服务器保持通讯
func serverProcessMes(conn net.Conn)  {
	//创建一个transfer实例，不停的读取服务器发送的消息
	tf:= &utils.Transfer{
		Conn: conn,
	}

	for{
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes,err := tf.ReadPkg()
		if err != nil{
			fmt.Println("tf.ReadPkg err=",err)
			return
		}
		//如果读取到消息，又是下一步处理逻辑
		fmt.Printf("mes=%v\n",mes)
	}
}

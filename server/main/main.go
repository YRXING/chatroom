package main

import (
	"chatroom/server/model"
	"fmt"
	"net"
	"time"
)

//处理客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//这里调用总控，创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯协程错误，err=", err)
		return
	}
}

//这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao(){
	//这里的pool本身就是一个全局变量，这里需要注意一下初始化顺序的问题
	//initPool后initUserdao
	model.MyUserDao = model.NewUserDao(pool)

}


func main() {
	//当服务启动时，我们就去初始化我们的redis链接池
	initPool("localhost:6379",16,0,300*time.Second)
	initUserDao()

	//提示信息
	fmt.Println("服务器在8889端口监听.....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//一旦监听成功，就等待库护短来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn, err := listen.Accept() //返回套接字
		if err != nil {
			fmt.Println("Listen.Accept err=", err)
			return
		}
		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}

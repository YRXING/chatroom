package process

import (
	"chatroom/client/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {
	//暂时不需要字段...
}

//给关联一个用户登录的方法
//写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2. 准备通过conn发送消息给服务
	mes := message.Message{
		Type: message.LoginMesType,
	}
	//3. 创建一个LoginMes结构体
	loginMes := message.LoginMes{
		UserPwd: userPwd,
		UserId:  userId,
	}
	//4. 将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5. 将data赋给mes.Data字段
	mes.Data = string(data)
	//6. 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//7. 这时候的data就是我们要发送的message消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	tf.WritePkg(data)
	fmt.Printf("客户端发送消息的长度=%d 内容=%s", len(data), string(data))

	//这里还需要处理服务器端返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		/*
			这里我们还需要在客户端启动一个协程，该协程保持和服务器端的通讯
			如果服务器有数据推送给客户端，则接受并显示在客户端的终端
		*/
		go serverProcessMes(conn)

		//1. 显示我们登录成功的菜单[循环显示]...
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}
	return
}

func (this *UserProcess) Register(userId int,userPwd string,userName string)(err error){
	//1. 连接到服务器
	conn,err:= net.Dial("tcp","localhost:8889")
	if err !=nil{
		fmt.Println("net.Dial err=",err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2. 准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3. 创建一个RegisterMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4. 将registerMes序列化
	data,err:=json.Marshal(registerMes)
	if err !=nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	//5. 把data赋给mes.Data字段
	mes.Data = string(data)
	//6. 将mes进行序列化
	data,err = json.Marshal(mes)
	if err !=nil{
		fmt.Println("json.Marshal err=",err)
		return
	}
	//创建一个transfer实例
	tf:= &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器
	err = tf.WritePkg(data)
	if err !=nil{
		fmt.Println("注册信息发送错误 err=",err)
		return
	}

	mes,err = tf.ReadPkg() //mes即时RegisterResMes
	if err !=nil{
		fmt.Println("readPkg(conn) err=",err)
		return
	}

	//将mes的Data部分反序列化成RegisterResMes
	var registerResMes  message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code ==200{
		fmt.Println("注册成功，现在可以登录了")
		os.Exit(0)
	}else{
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}
	return
}

package main

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

func login(userId int, userPwd string) (err error) {
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
	writePkg(conn,data)
	fmt.Printf("客户端发送消息的长度=%d 内容=%s", len(data), string(data))


	//这里还需要处理服务器端返回的消息
	mes,err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200{
		fmt.Println("登录成功")
	}else{
		fmt.Println(loginResMes.Error)
	}
	return
}

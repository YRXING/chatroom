package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

//处理客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//循环的客户端发送的消息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg(),返回message,err
		mes,err := readPkg(conn)
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出，服务器端也退出...")
				return
			}else{
				fmt.Println("readPkg err=",err)
				return
			}
		}
		err = serverProcessMes(conn,&mes)
		if err != nil{
			return
		}
	}
}

func readPkg(conn net.Conn)(mes message.Message,err error)  {
	buf := make([]byte,8096) //接收缓冲区
	fmt.Println("读取到客户端发送的数据...")
	//conn.Read在conn没有被关闭的情况下，才会阻塞，如果客户端关闭了，就不会阻塞
	_,err =conn.Read(buf[:4])	//先接收长度
	if err != nil{
		//err = errors.New("read pkg header error")
		return
	}

	//根据buf[:4]转成一个uint32类型
	pkgLen := binary.BigEndian.Uint32(buf[0:4])
	n,err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err !=nil{
		//err = errors.New("readk pkg bogy error")
		return
	}
	//把pkgLen反序列化成message.Message
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	return
}

func writePkg(conn net.Conn,data []byte)(err error)  {

	//1. 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)//将uint32转成byte切片
	//发送长度 conn.Write发送的是byte slice
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}


	//2. 发送data本身
	n,err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}

//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message)(err error)  {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
		err = serverProcessLogin(conn,mes)
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

//处理登录请求
func serverProcessLogin(conn net.Conn,mes *message.Message) (err error) {
	//1. 先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//2. 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//3. 在成名一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes
	//如果用户id=100,密码=123456，认为合法
	if loginMes.UserId == 100 && loginMes.UserPwd =="123456"{
		loginResMes.Code = 200
	}else{
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用..."
	}
	//4. 将loginResMes序列化
	data,err:= json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//5. 将data赋值给resMes
	resMes.Data = string(data)
	//6. 对resMes进行序列化，准备发送
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//7. 发送data,将其封装到writePkg中
	err = writePkg(conn,data)
	return
}

func main() {
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

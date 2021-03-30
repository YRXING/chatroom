package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

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
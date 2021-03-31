package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//这里将公共方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf  [8096]byte //传输缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取到客户端发送的数据...")
	//conn.Read在conn没有被关闭的情况下，才会阻塞，如果客户端关闭了，就不会阻塞
	_, err = this.Conn.Read(this.Buf[:4]) //先接收长度
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}

	//根据buf[:4]转成一个uint32类型
	pkgLen := binary.BigEndian.Uint32(this.Buf[0:4])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		//err = errors.New("readk pkg bogy error")
		return
	}
	//把pkgLen反序列化成message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	//1. 先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen) //将uint32转成byte切片
	//发送长度 conn.Write发送的是byte slice
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//2. 发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}

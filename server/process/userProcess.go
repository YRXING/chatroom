package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
	UserId int
}

//编写一个函数专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//1. 先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}
	//2. 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType
	//3. 在声明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes

	//我们需要到redis数据库去完成验证
	//使用model.MyUserDao到redis去验证
	user,err := model.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	if err!=nil{
		if err == model.ERROR_USER_NOTEXISTS{
			loginResMes.Code=500
			loginResMes.Error=err.Error()
		}else if err == model.ERROR_USER_PWD{
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		}else{
			loginResMes.Code=505
			loginResMes.Error="服务器内部错误..."
		}
	}else{
		loginResMes.Code=200
		//用户登录成功，我们就把该登录成功的人放入到userMgr中
		//将登录成功的用户userId赋给this
		this.UserId = loginMes.UserId
		//将当前在线用户的id放入到loginResMes.UserId,用于客户端展示
		userMgr.AddOnlineUser(this)
		//遍历userMgr.onlineUsers
		for id,_:=range userMgr.onlineUsers{
			loginResMes.UsersId = append(loginResMes.UsersId,id)
		}
		fmt.Println(user,"登录成功")
	}


	//4. 将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//5. 将data赋值给resMes
	resMes.Data = string(data)
	//6. 对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//7. 发送data,将其封装到writePkg中
	//因为使用分层模式（MVC），我们先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message)(err error)  {
	//先从mes中取出mes.Data,并直接反序列化成RegisterMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err!=nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}
	//1. 先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	var registerResMes message.RegisterResMes
	//2. 我们需要到redis数据库去完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err !=nil{
		if err ==model.ERROR_USER_EXISTS{
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()  //返回string
		}else{
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	}else {
		registerResMes.Code = 200
	}

	//3.对结果进行序列化
	data,err:=json.Marshal(registerResMes)
	if err!=nil{
		fmt.Println("json.Marshal fail",err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)
	//5. 对resMes进行序列化，准备发送
	data,err = json.Marshal(resMes)
	if err != nil{
		fmt.Println("json.Marshal fail",err)
		return
	}
	//6. 发送data
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}
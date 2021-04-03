package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMesType"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType		= "SmsMesType"
)

//这里我们定义几个用户状态的常量
const(
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //序列化后的消息体
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态码 500 表示用户未注册 200表示登录成功
	Error string `json:"error"` //返回错误信息
	UsersId []int 	`json:"usersId"`
}

type RegisterMes struct{
	User User `json:"user"` //类型就是User结构体
}

type RegisterResMes struct {
	Code int `json:"code"` //返回状态码
	Error string `json:"error"` //返回错误信息
}

//为了配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

//发送的消息
type SmsMes struct {
	Content string `json:"content"` //内容
	User //匿名结构体，继承
}
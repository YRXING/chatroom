package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMesType"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //序列化后的消息体
}

//定义两种消息..后面需要在增加

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
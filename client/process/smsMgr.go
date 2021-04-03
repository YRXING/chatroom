package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outputGroupMes(mes *message.Message){//这个地方一定是smsMes
	//显示即可
	//1 反序列化mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err!=nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	//显示信息
	info := fmt.Sprintf("用户id:\t%d：\t%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}

package process

import "fmt"

/*
UserMgr实例在服务器端有且只有一个
因为在很多地方都会使用到，因此我们将其定义为全局变量
*/

var(
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

//完成对userMgr初始化工作
func init(){
	userMgr = &UserMgr{
		onlineUsers: make(map[int]*UserProcess,1024),
	}
}
//完成对onlineUsers添加
func (this *UserMgr) AddOnlineUser(up *UserProcess){
	this.onlineUsers[up.UserId] = up
}

//删除
func (this *UserMgr) DelOnlineUser(userId int)  {
	delete(this.onlineUsers,userId)
}

//返回当前所有在线的用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess  {
	return this.onlineUsers
}

//根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int)(up *UserProcess,err error)  {
	//如何从map取出一个值，带检测方式
	up,ok := this.onlineUsers[userId] //多返回值默认返回第一个
	if !ok{ //说明，你要查找的这个用户，当前不在线
		err = fmt.Errorf("用户%d 不存在",userId)
		return
	}
	return
}


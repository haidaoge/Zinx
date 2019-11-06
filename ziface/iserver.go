/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-24 13:48:54
 * @LastEditTime: 2019-09-24 13:48:54
 * @LastEditors: your name
 */
package ziface

//定义服务器接口
type IServer interface {
	//启动服务器方法
	Start()
	//停止服务器方法
	Stop()
	//开启业务服务方法
	Serve()
	//路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(msgId uint32, router IRouter)
}

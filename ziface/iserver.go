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
}

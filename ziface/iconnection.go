/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-15 15:01:04
 * @LastEditTime: 2019-10-15 15:15:32
 * @LastEditors: Please set LastEditors
 */
package ziface

import "net"

// 定义接口
type IConnection interface {
	// 启动连接，让当前连接开始工作
	Start()
	// 停止连接，结束当前连接状态
	Stop()
	// 从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	// 获取当前IP
	GetConnID() uint32
	// 获取远程客户端地址信息
	RemoteAddr() net.Addr
}

//定义一个统一处理链接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-15 17:00:52
 * @LastEditTime: 2019-10-15 17:00:52
 * @LastEditors: your name
 */
package main

import (
	"zinx/znet"
)

//Server 模块的测试函数
func main() {

	//1 创建一个server 句柄 s
	s := znet.NewServer("[zinx V0.1]")

	//2 开启服务
	s.Serve()
}

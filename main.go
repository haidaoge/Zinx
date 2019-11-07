package main

import (
	"fmt"
	"zinx/routers"
	"zinx/ziface"
	"zinx/znet"
)

func main() {
	//创建一个server句柄
	s := znet.NewServer()

	//注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	routers.InitRouter(s)

	//开启服务
	s.Serve()
}

//创建连接的时候执行
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnecionBegin is Called ... ")

	//=============设置两个链接属性，在连接创建之后===========
	fmt.Println("Set conn Name, Home done!")
	conn.SetProperty("Name", "Aceld")
	conn.SetProperty("Home", "https://www.jianshu.com/u/35261429b7f1")
	//===================================================

	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
}

//连接断开的时候执行
func DoConnectionLost(conn ziface.IConnection) {
	//============在连接销毁之前，查询conn的Name，Home属性=====
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Conn Property Home = ", home)
	}
	//===================================================

	fmt.Println("DoConneciotnLost is Called ... ")
}

package routers

import (
	"zinx/controllers/example"
	"zinx/ziface"
)

func InitRouter(s ziface.IServer) {
	//配置路由
	s.AddRouter(0, &example.PingRouter{})
	s.AddRouter(1, &example.HelloZinxRouter{})
}

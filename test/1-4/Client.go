/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-10-15 17:03:47
 * @LastEditTime: 2019-10-15 17:03:47
 * @LastEditors: your name
 */
package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	fmt.Println("Client Test ... start")
	//3秒之后发起测试请求，给服务端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:6666")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		_, err := conn.Write([]byte("客户端消息"))
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error ")
			return
		}

		fmt.Printf(" server call back : %s, cnt = %d\n", buf, cnt)

		time.Sleep(1 * time.Second)
	}
}

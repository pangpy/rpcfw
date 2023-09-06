package main

import (
	"log"
	"net"
	"net/rpc"
)

// 服务端结构
type Calculator struct{}

// 计算两个整数的和
func (c *Calculator) Add(args *Args, result *int) error {
	*result = args.A + args.B
	return nil
}

// 输入参数结构
type Args struct {
	A, B int
}

func main() {
	calculator := new(Calculator)
	rpc.Register(calculator)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("监听端口失败：", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("建立连接失败：", err)
		}
		go rpc.ServeConn(conn)
	}
}

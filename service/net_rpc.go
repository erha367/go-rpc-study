package service

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2020/12/3 下午7:09
 * @company    ：eeo.cn
 */

// 算数运算请求结构体
type ArithRequest struct {
	A int
	B int
}

// 算数运算响应结构体
type ArithResponse struct {
	Pro int // 乘积
}

type IArith interface {
	Multiply(req ArithRequest, res *ArithResponse) error
}

type Arith struct{}

// 乘法运算方法
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {
	res.Pro = req.A * req.B
	return nil
}

func NetRpc(){
	rpc.Register(new(Arith)) // 注册rpc服务
	rpc.HandleHTTP()                 // 采用http协议作为rpc载体
	lis, err := net.Listen("tcp", "127.0.0.1:8095")
	if err != nil {
		log.Fatalln(`rpc err `, err)
	}
	http.Serve(lis, nil)
}

//
func JsonRpc() {
	rpc.Register(new(Arith))
	lis, err := net.Listen("tcp", "127.0.0.1:8096")
	if err != nil {
		log.Println(`tcp listen err `, err)
		return
	}
	for {
		conn, err := lis.Accept() // 接收客户端连接请求
		if err != nil {
			log.Println(`for json rpc err `, err)
			continue
		}
		go func(conn net.Conn) { // 并发处理客户端请求
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}

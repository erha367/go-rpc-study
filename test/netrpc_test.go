package test

import (
	"context"
	"google.golang.org/grpc"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpcx/proto"
	"rpcx/service"
	"testing"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2020/12/3 下午7:14
 * @company    ：eeo.cn
 */

func TestNetrpc(t *testing.T) {
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8095")
	if err != nil {
		t.Error(`http dial err`, err)
	}
	//调用乘法
	req := service.ArithRequest{9, 2}
	var res service.ArithResponse
	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
	if err != nil {
		t.Error("arith error: ", err)
	}
	t.Log(res.Pro)
	//调用加法
	req2 := service.ArithRequest{
		A: 50,
		B: 60,
	}
	var res2 service.ArithResponse
	err = conn.Call("Jia.Add", req2, &res2)
	if err != nil {
		t.Log(err)
	}
	t.Log(res2.Pro)
	defer conn.Close()
}

func TestJsonrpc(t *testing.T) {
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8096")
	if err != nil {
		t.Error(err)
	}
	req := service.ArithRequest{9, 20}
	var res service.ArithResponse
	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
	if err != nil {
		t.Error("arith error: ", err)
	}
	t.Log(res.Pro)
	defer conn.Close()
}

func TestGrpc(t *testing.T) {
	conn, err := grpc.Dial(`127.0.0.1:8028`, grpc.WithInsecure())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	client := proto.NewArithServiceClient(conn)
	ret, err := client.Multiply(context.Background(), &proto.ArithRequest{
		A: 3,
		B: 600,
	})
	t.Log(ret, err)
}

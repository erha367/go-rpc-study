package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpcx/consul"
	"rpcx/proto"
	"rpcx/service"
	"testing"
	"time"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2020/12/3 下午7:14
 * @company    ：eeo.cn
 */

func TestNetrpc(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
		t.Log(`这里写继续的业务`)
	}()
	h := `127.0.0.1:8095`
	conn, err := rpc.DialHTTP("tcp", h)
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
	time.Sleep(time.Second * 5)
	t.Log(`业务的正常逻辑，啊哈哈哈`)
}

func RpcDial() error {
	h := `127.0.0.1:8095`
	conn, err := rpc.DialHTTP("tcp", h)
	if err != nil {
		return err
	}
	//调用乘法
	req := service.ArithRequest{9, 2}
	var res service.ArithResponse
	err = conn.Call("Arith.Multiply", req, &res) // 乘法运算
	if err != nil {
		return err
	}
	log.Println(res.Pro)
	return nil
}

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
		t.Log(`这里写继续的业务`)
	}()
	RpcDial()
	t.Log(`这里是原来的业务逻辑`)
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
	schema, err := consul.GenerateAndRegisterConsulResolver(`consul.sit13.dom:61170`, `yangsen`)
	if err != nil {
		log.Fatal("init consul resovler err", err.Error())
	}

	//建立连接
	conn, err := grpc.Dial(fmt.Sprintf("%s:///yangsen", schema), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewHelloServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	res, err := client.SayHello(ctx, &proto.HelloRequest{
		Name:                 "152****4335",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	log.Println(res, err)
	c2 := proto.NewClassServiceClient(conn)
	res2, err := c2.Enter(ctx, &proto.EnterReq{
		Account:              "2",
		AccountType:          0,
		Nickname:             "",
		Cookie:               "",
		Source:               0,
		Client:               0,
		LessonKey:            "",
		ClassId:              0,
		Ip:                   "127.0.0.1",
		HasCookie:            false,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	log.Println(res2, err)
}

/*
func TestGrpc2(t *testing.T) {
	conn, err := grpc.Dial(`127.0.0.1:9000`, grpc.WithInsecure())
	if err != nil {
		t.Error(`rpc dial err`, err)
	}
	defer conn.Close()
	client := proto.NewDemoClient(conn)
	res, err := client.SayHelloURL(context.Background(), &proto.HelloReq{
		Name:                 "Ma da Ha",
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	})
	t.Log(res, err)
}
*/

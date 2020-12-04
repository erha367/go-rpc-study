package main

import (
	"log"
	"os"
	"os/signal"
	"rpcx/service"
	"syscall"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2020/12/3 下午7:08
 * @company    ：eeo.cn
 */

func main() {
	sigc := make(chan os.Signal)
	/** grpc **/
	go service.GrpcInit()
	/** jsonrpc **/
	go service.JsonRpc()
	/** net/rpc库 **/
	go service.NetRpc()
	/*-  系统信号监听 -*/
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	sig := <- sigc
	log.Println(sig)
}

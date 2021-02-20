# go 语言rpc学习笔记
## 代码结构
```
go-rpc-study/
├── LICENSE
├── README.md
├── consul      //服务注册、发现
│   ├── consul_reg.go
│   └── consul_resolver.go
├── go.mod
├── go.sum
├── main.go     //主程序入口
├── proto       //proto文件及生成的go代码
│   ├── class.pb.go
│   ├── class.proto
│   ├── hello.bm.go
│   ├── hello.pb.go
│   ├── hello.proto
│   └── hello.swagger.json
├── service     //定义服务
│   ├── grpc.go
│   └── net_rpc.go
└── test        //测试用例
    ├── curl_test.go
    ├── jsonrpc_client.php
    └── netrpc_test.go
```

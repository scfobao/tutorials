package main

import (
	"fmt"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/auth/handler"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/auth/model"
	s "github.com/micro-in-cn/tutorials/microservice-in-micro/part5/auth/proto/auth"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/basic/common"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part5/plugins/graylog"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-plugins/config/source/grpc"
	"github.com/smallnest/rpcx/log"
)

var (
	appName = "auth_srv"
	cfg     = &authCfg{}
)

type authCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置、数据库等信息
	initCfg()
	graylog.GetLog(cfg).Write(map[string]interface{}{"data": "hello i am auth srv"})
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name(cfg.Name),
		micro.Registry(micReg),
		micro.Version(cfg.Version),
		micro.Address(cfg.Addr()),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterServiceHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Error("[main] error")
		panic(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := &common.Consul{}
	err := config.C().App("consul", consulCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.Host, consulCfg.Port)}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(
		config.WithSource(source),
		config.WithApp(appName),
	)

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	return
}

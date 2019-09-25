package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part6/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part6/basic/common"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part6/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part6/payment-web/handler"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"github.com/micro/go-plugins/config/source/grpc"
)

var (
	appName             = "payment_web"
	cfg                 = &appCfg{}
	configServerAddress = "192.168.1.232:9600"
)

type appCfg struct {
	common.AppCfg
}

func main() {
	// 初始化配置
	initCfg()
	configServerAddress = os.Getenv("CONFIG_SERVER") + ":" + os.Getenv("CONFIG_SERVER_PORT")
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)

	// 创建新服务
	service := web.NewService(
		web.Name(cfg.Name),
		web.Version(cfg.Version),
		web.RegisterTTL(time.Second*15),
		web.RegisterInterval(time.Second*10),
		web.Registry(micReg),
		web.Address(cfg.Addr()),
	)

	// 初始化服务
	if err := service.Init(
		web.Action(
			func(c *cli.Context) {
				// 初始化handler
				handler.Init()
			}),
	); err != nil {
		log.Fatal(err)
	}

	// 新建订单接口
	authHandler := http.HandlerFunc(handler.PayOrder)
	service.Handle("/payment/pay-order", handler.AuthWrapper(authHandler))

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
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
		grpc.WithAddress(configServerAddress),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg] 配置，cfg：%v", cfg)

	return
}

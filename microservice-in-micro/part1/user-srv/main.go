package main

import (
	"fmt"
	"time"

	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/basic"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/basic/config"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/handler"
	"github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/model"
	s "github.com/micro-in-cn/tutorials/microservice-in-micro/part1/user-srv/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()
	log.Logf("[main] 初始化成功")

	// 使用etcd注册
	micReg := consul.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.user"),
		micro.RegisterTTL(time.Second*15),
		micro.RegisterInterval(time.Second*10),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// 注册服务
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()

	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}

}

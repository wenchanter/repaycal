package main

import (
	"flag"
	"fmt"

	"repaycal/apps/calculator/rpc/internal/config"
	"repaycal/apps/calculator/rpc/internal/server"
	"repaycal/apps/calculator/rpc/internal/svc"
	"repaycal/apps/calculator/rpc/protoc/calculator"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

var configFile = flag.String("f", "etc/calculator.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		calculator.RegisterCalculateServer(grpcServer, server.NewCalculateServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

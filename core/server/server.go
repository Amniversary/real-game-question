package server

import (
	"time"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/Amniversary/real-game-question/core/logger"
)

type Server interface {
	Init()
	Service() micro.Service
	Run() error
}

func NewServer(opt ...Option) Server {
	return newMicroServer(opt...)
}

type microServer struct {
	opts Options

	service micro.Service
}

func newMicroServer(opts ...Option) Server {
	options := newOptions(opts...)

	return &microServer{
		opts: options,
	}
}

func (s *microServer) Init() {
	r := etcdv3.NewRegistry(registry.Addrs(s.opts.Address...))

	s.service = micro.NewService(
		// server basic info
		// todo: 服务基础信息
		micro.Name(s.opts.ServerName),
		micro.Version(s.opts.Version),
		// register of etcd
		// todo: 注册 Etcd
		micro.RegisterTTL(
			time.Duration(s.opts.RegisterTTL)*time.Second,
		),
		micro.RegisterInterval(
			time.Duration(s.opts.RegisterInterval)*time.Second,
		),
		micro.Registry(r),
		// todo: 日志注册
		// infrastructure of log
		// TODO: need ELK
		micro.WrapHandler(logger.ServerLogWrapper),
		micro.WrapClient(logger.ClientLogWrap),
	)

	//s.service.Init()
}

func (s *microServer) Run() (err error) {
	if err = s.service.Run(); err != nil {
		log.Fatalf("service run error: %v", err)
	}
	return
}

func (s *microServer) Service() micro.Service {
	return s.service
}

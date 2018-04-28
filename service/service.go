package service

import (
	"github.com/Amniversary/real-game-question/config"
	"github.com/Amniversary/real-game-question/core/server"
	"github.com/Amniversary/real-game-question/models"
	proto "github.com/Amniversary/real-game-question/proto"
)

func Run(c *config.Config) {
	s := server.NewServer(
		server.ServerName(c.ServerName),
		server.Version(c.Version),
		server.EtcdAddress(c.Etcd.Address),
		server.EtcdRegisterTTL(c.Etcd.RegisterTTL),
		server.EtcdRegisterInterval(c.Etcd.RegisterInterval),
	)
	s.Init()
	// todo: init database
	models.NewMysql(c)

	proto.RegisterQuestionHandler(s.Service().Server(), &Question{Client:s.Service().Client()})

	s.Run()
}

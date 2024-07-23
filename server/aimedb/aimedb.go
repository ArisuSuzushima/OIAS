package aimedb

import (
	"github.com/panjf2000/gnet"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*gnet.EventServer
}

func (s *Server) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	logrus.Info(
		"AimeDB server is listening on: ",
		srv.Addr.String(),
	)
	return
}

func (s *Server) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func Run() {
	server := &Server{}

	err := gnet.Serve(server, "tcp://:22356", gnet.WithMulticore(true), gnet.WithReusePort(true))
	if err != nil {
		logrus.Fatalf("AiMeDB server error: %v", err)
	}
}

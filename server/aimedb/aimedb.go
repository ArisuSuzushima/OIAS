package aimedb

import (
	"github.com/panjf2000/gnet"
	"github.com/sirupsen/logrus"
)

type Server struct {
	*gnet.EventServer
}

func (s *Server) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	logrus.Info("AimeDB server is listening on: ", srv.Addr.String())
	return
}

func (s *Server) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	header, err := DecodeHeader(frame)
	if err != nil {
		logrus.Error("Failed to decode header: ", err)
		return
	}

	logrus.Infof("Received packet with CommandID: %d", header.CommandID)

	responseHeader := &AimeDbHeader{
		Magic:     magic,
		Version:   version,
		CommandID: header.CommandID,
		Length:    headerSize,
		Result:    0,
	}
	response, err := EncodeHeader(responseHeader)
	if err != nil {
		logrus.Error("Failed to encode response: ", err)
		return
	}

	return response, gnet.None
}

func Run() {
	server := &Server{}
	err := gnet.Serve(server, "tcp://:22356", gnet.WithMulticore(true), gnet.WithReusePort(true))
	if err != nil {
		logrus.Fatalf("AimeDB server error: %v", err)
	}
}

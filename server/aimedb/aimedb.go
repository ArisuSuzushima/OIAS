package aimedb

import (
	"bytes"
	"github.com/panjf2000/gnet"
	"github.com/sirupsen/logrus"
	"sync"
)

type Server struct {
	*gnet.EventServer
	handlers map[uint16]CommandHandler // Command handlers
	mu       sync.Mutex
	buffers  map[gnet.Conn]*bytes.Buffer // Buffers for each connection
}

// NewServer creates a new AimeDB server
func NewServer() *Server {
	return &Server{
		handlers: make(map[uint16]CommandHandler),
		buffers:  make(map[gnet.Conn]*bytes.Buffer),
	}
}

// RegisterHandler registers a command handler
func (s *Server) RegisterHandler(cmdID uint16, handler CommandHandler) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[cmdID] = handler
}

func (s *Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	s.buffers[c] = &bytes.Buffer{}
	return
}

func (s *Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	delete(s.buffers, c)
	return
}

func (s *Server) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	buf := s.buffers[c]
	buf.Write(frame)

	for {
		if buf.Len() < HeaderSize {
			return
		}

		header, err := DecodeHeader(buf.Bytes()[:HeaderSize])
		if err != nil {
			logrus.Error("Header decode failed:", err)
			return nil, gnet.Close
		}

		fullPacketSize := HeaderSize + int(header.Length)
		if buf.Len() < fullPacketSize {
			return // Wait for more data
		}

		encryptedData := buf.Next(fullPacketSize)
		s.handleRequest(c, encryptedData, header)
	}
}

func (s *Server) handleRequest(c gnet.Conn, encryptedData []byte, header *AimeDbHeader) {
	//
	decryptedFull, err := Decrypt(encryptedData)
	if err != nil {
		logrus.Error("Full decrypt failed:", err)
		return
	}

	payload := decryptedFull[HeaderSize:]
	handler, exists := s.handlers[header.CommandID]
	if !exists {
		logrus.Warnf("Unhandled CommandID: 0x%04X", header.CommandID)
		s.sendErrorResponse(c, header, 0xFFFF) // Unsupported command
		return
	}
	resultCode, responsePayload, err := handler.Handle(header, payload)
	if err != nil {
		logrus.Errorf("Handler failed for 0x%04X: %v", header.CommandID, err)
		s.sendErrorResponse(c, header, resultCode)
		return
	}

	respHeader := &AimeDbHeader{
		Magic:     MagicValue,
		Version:   VersionValue,
		CommandID: header.CommandID,
		Length:    uint16(len(responsePayload)),
		Result:    resultCode,
		GameID:    header.GameID,
		StoreID:   header.StoreID,
		KeychipID: header.KeychipID,
	}

	encryptedResp, err := EncodeResponse(respHeader, responsePayload)
	if err != nil {
		logrus.Error("Encode response failed:", err)
		return
	}
	_ = c.AsyncWrite(encryptedResp)
}

func (s *Server) sendErrorResponse(c gnet.Conn, origHeader *AimeDbHeader, resultCode uint16) {
	errHeader := &AimeDbHeader{
		Magic:     MagicValue,
		Version:   VersionValue,
		CommandID: origHeader.CommandID,
		Length:    0,
		Result:    resultCode,
		GameID:    origHeader.GameID,
		StoreID:   origHeader.StoreID,
		KeychipID: origHeader.KeychipID,
	}
	encryptedResp, _ := EncodeResponse(errHeader, nil)
	_ = c.AsyncWrite(encryptedResp)
}

package handlers

import (
	"OMG_ITS_ALLNET_SERVER/server/aimedb"
	"OMG_ITS_ALLNET_SERVER/utils"
	"fmt"
)

type SystemHandler struct{}

func (h *SystemHandler) Handle(header *aimedb.AimeDbHeader, payload []byte) (uint16, []byte, error) {
	if header.CommandID == utils.DecToHex(100) {
		// TODO: Implement
		return 0x0000, []byte{0x01}, nil // Example success response
	}
	return 0xFFFF, nil, fmt.Errorf("unsupported system command")
}

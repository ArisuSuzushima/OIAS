package handlers

import (
	"OMG_ITS_ALLNET_SERVER/server/aimedb"
	"OMG_ITS_ALLNET_SERVER/utils"
)

type FeliCaHandler struct{}

func (h *FeliCaHandler) Handle(header *aimedb.AimeDbHeader, payload []byte) (uint16, []byte, error) {
	switch utils.HexToDec(header.CommandID) {
	case 1:
		if len(payload) < 16 {
			return 0, nil, nil
		} else {
			idm := payload[:8]
			pmm := payload[8:]
			return 0, GetFelicaID(pmm, idm), nil
		}
	case 2:
		break
	case 3:
		break
	}

	return 0, nil, nil
}

func GetFelicaID(pmm []byte, idm []byte) []byte {
	return nil
}

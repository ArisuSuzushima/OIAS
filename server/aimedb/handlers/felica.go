package handlers

import (
	"OMG_ITS_ALLNET_SERVER/server/aimedb"
	"OMG_ITS_ALLNET_SERVER/utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
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

// GetFelicaID calculate access code by slicing idm
// Returns []byte(ac) AccessCode converted to Byte
func GetFelicaID(pmm []byte, idm []byte) []byte {
	idmStr := string(idm)

	// Handle empty card idm
	if idmStr == "00000000000000000000" {
		// TODO: detect which client sent an empty id
		logrus.Debug("FeliCaHandler::GetFelicaID Receive empty idm!")
		// TODO: exec safety action
		return []byte("00000000000000000000")
	}

	// Convert idm from Hex to D
	hexValue, err := strconv.ParseInt(idmStr, 16, 64)
	if err != nil {
		logrus.Debug("FeliCaHandler::GetFelicaID FelicaID parse error:", err)
		return []byte("00000000000000000000")
	}

	// Format to 20-bit Hex, prefixed with 0
	ac := fmt.Sprintf("%020x", hexValue)
	return []byte(ac)
}

func RegFelicaID(pmm []byte, idm []byte) []byte {
	return nil
}

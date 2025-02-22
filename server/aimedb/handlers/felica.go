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
			return 0, GetFeliCaID(idm, pmm), nil
		}
	case 2:
		break
	case 3:
		break
	}

	return 0, nil, nil
}

// GetFeliCaID calculate access code by slicing idm
// Returns []byte(ac) AccessCode converted to Byte
func GetFeliCaID(idm []byte, pmm []byte) []byte {
	idmStr := string(idm)

	// Handle empty card idm
	if idmStr == "00000000000000000000" {
		// TODO: detect which client sent an empty id
		logrus.Debug("FeliCaHandler::GetFeliCaID Receive empty idm!")
		// TODO: exec safety action
		return []byte("00000000000000000000")
	}

	// Convert idm from Hex to D
	hexValue, err := strconv.ParseInt(idmStr, 16, 64)
	if err != nil {
		logrus.Debug("FeliCaHandler::GetFeliCaID FeliCaID parse error:", err)
		return []byte("00000000000000000000")
	}

	// Format to 20-bit Hex, prefixed with 0
	ac := fmt.Sprintf("%020x", hexValue)
	logrus.Debug("FeliCaHandler::GetFeliCaID Rec/Gen FeliCaID:", ac)
	return []byte(ac)
}

func RegFeliCaID(idm []byte, pmm []byte) []byte {
	return nil
}

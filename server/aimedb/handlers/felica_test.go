package handlers

import (
	"testing"
)

func TestGetFeliCaID(t *testing.T) {
	idm := "01020123456789AB"
	t.Log("FeliCaHandler::TestGetFeliCaID Set IDM:", idm)
	pmm := "A1B2C3D4E5F60708"
	t.Log("FeliCaHandler::TestGetFeliCaID Set PMM:", pmm)
	acc := string(GetFeliCaID([]byte(idm), []byte(pmm)))
	if len(acc) != 20 {
		t.Fatal("FeliCaHandler::TestGetFeliCaID Broken ACC:", acc)
	}
	t.Log("FeliCaHandler::TestGetFeliCaID Access Code:", acc)
}

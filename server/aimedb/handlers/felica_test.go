package handlers

import (
	"testing"
)

func TestGetFelicaID(t *testing.T) {
	idm := []byte("01020123456789AB")
	pmm := []byte("A1B2C3D4E5F60708")
	acc := string(GetFelicaID(pmm, idm))
	t.Log("FeliCaHandler::TestGetFelicaID Access Code:", acc)
}

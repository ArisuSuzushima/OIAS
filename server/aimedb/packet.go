package aimedb

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const headerSize = 32

// AimeDbHeader represents the packet header
type AimeDbHeader struct {
	Magic     uint16
	Version   uint16
	CommandID uint16
	Length    uint16
	Result    uint16
	GameID    [6]byte
	StoreID   [4]byte
	KeychipID [12]byte
}

const (
	magic   uint16 = 0xa13e
	version uint16 = 0x3087
)

func DecodeHeader(data []byte) (*AimeDbHeader, error) {
	decrypted, err := Decrypt(data)
	if err != nil {
		return nil, err
	}
	header := &AimeDbHeader{}
	err = binary.Read(bytes.NewReader(decrypted[:headerSize]), binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}
	if header.Magic != magic {
		return nil, fmt.Errorf("invalid magic: %x", header.Magic)
	}
	return header, nil
}

func EncodeHeader(header *AimeDbHeader) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, header)
	if err != nil {
		return nil, err
	}
	encrypted, err := Encrypt(buf.Bytes())
	if err != nil {
		return nil, err
	}
	return encrypted, nil
}

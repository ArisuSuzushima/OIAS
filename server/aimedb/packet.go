package aimedb

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	MagicValue   = 0xA13E
	VersionValue = 0x3087
	HeaderSize   = 32 // All headers are 32 bytes
)

// AimeDbHeader represents the header of an AimeDB packet
type AimeDbHeader struct {
	Magic     uint16   // Magic 0xA13E
	Version   uint16   // Version 0x3087
	CommandID uint16   // Command ID
	Length    uint16   // Length of the payload
	Result    uint16   // Result
	_         [2]byte  // Idk what this is(?)
	GameID    [6]byte  // Game ID
	StoreID   [4]byte  // Tenpo ID
	KeychipID [12]byte // Keychip ID
}

// DecodeHeader decodes the header of an AimeDB packet
func DecodeHeader(encryptedData []byte) (*AimeDbHeader, error) {
	decrypted, err := Decrypt(encryptedData[:HeaderSize])
	if err != nil {
		return nil, fmt.Errorf("decrypt failed: %w", err)
	}

	header := &AimeDbHeader{}
	reader := bytes.NewReader(decrypted)
	if err := binary.Read(reader, binary.LittleEndian, header); err != nil {
		return nil, fmt.Errorf("binary read failed: %w", err)
	}

	if header.Magic != MagicValue {
		return nil, fmt.Errorf("invalid magic: 0x%X", header.Magic)
	}
	if header.Version != VersionValue {
		return nil, fmt.Errorf("invalid version: 0x%X", header.Version)
	}

	return header, nil
}

// EncodeResponse encodes a response packet
func EncodeResponse(header *AimeDbHeader, payload []byte) ([]byte, error) {
	buf := new(bytes.Buffer)

	header.Length = uint16(len(payload))

	if err := binary.Write(buf, binary.LittleEndian, header); err != nil {
		return nil, err
	}

	fullData := append(buf.Bytes(), payload...)
	encrypted, err := Encrypt(fullData)
	if err != nil {
		return nil, err
	}

	return encrypted, nil
}

package internal

import (
	"bytes"
	"encoding/binary"
)

func ReadPacket(buf []byte, pack interface{}) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, pack); err != nil {
		return err
	}

	return nil
}

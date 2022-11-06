package udp

import (
	"bytes"
	"encoding/binary"
)

func ReadPacket(buf []byte, pack interface{}) error {
	reader := bytes.NewReader(buf)
	return binary.Read(reader, binary.LittleEndian, pack)
}

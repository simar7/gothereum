package main

import "encoding/binary"

func packUint16(data string, size uint16) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint16(buf, size)
}

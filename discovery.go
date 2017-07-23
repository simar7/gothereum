package main

import (
	"log"
	"net"
)

type EndPoint struct {
	addr    net.IPAddr
	udpPort net.UDPAddr
	tcpPort net.TCPAddr
}

func (e *EndPoint) pack() EndPoint {
	// FIXME: data types for packing
	return EndPoint{
		addr: e.addr,
		//udpPort: packUint16(e.udpPort.Port, uint16(2)),
		//tcpPort: packUint16(e.tcpPort.Port, uint16(2)),
	}
}

type PingNode struct {
	version    uint64
	packetType uint64
	srcURL     string
	dstURL     string
}

func (p *PingNode) new(srcURL string, dstURL string) PingNode {
	return PingNode{
		version:    Version,
		packetType: PacketType,
		srcURL:     srcURL,
		dstURL:     dstURL,
	}
}

func (p *PingNode) pack() PingNode {
	// TODO: Return a new populated PingNode object
	return PingNode{}
}

func main() {
	log.Print("This is a WIP, please check back later!")
}

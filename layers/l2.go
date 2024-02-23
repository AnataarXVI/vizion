package layers

import (
	"encoding/binary"
	"net"

	"github.com/AnataarXVI/vizion/buffer"
)

// Ethernet Types
var ETHERTYPE = map[uint16]string{
	0x0800: "IPv4",
	0x0806: "ARP",
	0x86dd: "IPv6",
}

type Ether struct {
	Dst  net.HardwareAddr
	Src  net.HardwareAddr
	Type uint16
}

// Create and return an Ether layer with default value set
func EtherLayer() Ether {
	eth := Ether{}
	eth.SetDefault()
	return eth
}

// GetName returns the protocol name.
func (e *Ether) GetName() string {
	return "Ethernet"
}

// Set a default value for each layer field
func (e *Ether) SetDefault() {

	ifaces, _ := net.Interfaces()

	e.Dst = net.HardwareAddr{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}
	e.Src = ifaces[1].HardwareAddr
	e.Type = 0x0900
}

// Serialize will convert a layer into bytes
func (e *Ether) Build() *buffer.ProtoBuff {

	// Initiate the buffer
	var buffer buffer.ProtoBuff

	// Write data field into the buffer
	buffer.Add("Dst", e.Dst, nil)

	buffer.Add("Src", e.Src, nil)

	buffer.Add("Type", e.Type, ETHERTYPE[e.Type])

	return &buffer
}

// Deserialize will convert bytes into a layer
func (e *Ether) Dissect(buffer *buffer.ProtoBuff) *buffer.ProtoBuff {
	e.Dst = buffer.Next(6)
	e.Src = buffer.Next(6)
	e.Type = binary.BigEndian.Uint16(buffer.Next(2))
	return buffer
}

// BindLayer return the top
func (e *Ether) BindLayer() Layer {

	// If ARP
	if e.Type == 0x0806 {
		return &ARP{}
	}

	return nil
}

var HARDWARE_TYPES = map[uint16]string{
	1:  "Ethernet (10Mb)",
	2:  "Ethernet (3Mb)",
	3:  "AX.25",
	4:  "Proteon ProNET Token Ring",
	5:  "Chaos",
	6:  "IEEE 802 Networks",
	7:  "ARCNET",
	8:  "Hyperchannel",
	9:  "Lanstar",
	10: "Autonet Short Address",
	11: "LocalTalk",
	12: "LocalNet",
	13: "Ultra link",
	14: "SMDS",
	15: "Frame relay",
	16: "ATM",
	17: "HDLC",
	18: "Fibre Channel",
	19: "ATM",
	20: "Serial Line",
	21: "ATM",
}

var OPCODE = map[uint16]string{
	1: "who-has",
	2: "is-at",
	3: "RARP-req",
	4: "RARP-rep",
	5: "Dyn-RARP-req",
	6: "Dyn-RAR-rep",
	7: "Dyn-RARP-err",
	8: "InARP-req",
	9: "InARP-rep",
}

type ARP struct {
	Hwtype uint16
	Ptype  uint16
	Hwlen  uint8
	Plen   uint8
	Opcode uint16
	Hwsrc  net.HardwareAddr
	Psrc   net.IP
	Hwdst  net.HardwareAddr
	Pdst   net.IP
}

func (a *ARP) GetName() string {
	return "ARP"
}

// Create and return an ARP layer with default value set
func ARPLayer() ARP {
	arp := ARP{}
	arp.SetDefault()
	return arp
}

// Set a default value for each layer field
func (a *ARP) SetDefault() {

	ifaces, _ := net.Interfaces()
	netAddr, _ := net.InterfaceAddrs()

	a.Hwtype = 1
	a.Ptype = 0x0800
	a.Hwlen = 6
	a.Plen = 4
	a.Opcode = 1
	a.Hwsrc = ifaces[1].HardwareAddr
	a.Psrc = netAddr[1].(*net.IPNet).IP.To4()
	a.Hwdst = net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	a.Pdst = net.IP{0, 0, 0, 0}

}

// TODO: Take into account the addition of paddind depending on frame size
func (a *ARP) Build() *buffer.ProtoBuff {
	// Initiate the buffer
	var buffer buffer.ProtoBuff

	buffer.Add("Hwtype", a.Hwtype, HARDWARE_TYPES[a.Hwtype])

	buffer.Add("Ptype", a.Ptype, ETHERTYPE[a.Ptype])

	buffer.Add("Hwlen", a.Hwlen, nil)

	buffer.Add("Plen", a.Plen, nil)

	buffer.Add("Opcode", a.Opcode, OPCODE[a.Opcode])

	buffer.Add("Hwsrc", a.Hwsrc, nil)

	buffer.Add("Psrc", a.Psrc, nil)

	buffer.Add("Hwdst", a.Hwdst, nil)

	buffer.Add("Pdst", a.Pdst, nil)

	return &buffer
}

func (a *ARP) Dissect(buffer *buffer.ProtoBuff) *buffer.ProtoBuff {
	a.Hwtype = binary.BigEndian.Uint16(buffer.Next(2))
	a.Ptype = binary.BigEndian.Uint16(buffer.Next(2))
	a.Hwlen = uint8(buffer.Next(1)[0])
	a.Plen = uint8(buffer.Next(1)[0])
	a.Opcode = binary.BigEndian.Uint16(buffer.Next(2))
	a.Hwsrc = buffer.Next(int(a.Hwlen))
	a.Psrc = buffer.Next(int(a.Plen))
	a.Hwdst = buffer.Next(int(a.Hwlen))
	a.Pdst = buffer.Next(int(a.Plen))
	return buffer
}

func (a *ARP) BindLayer() Layer {
	return nil
}

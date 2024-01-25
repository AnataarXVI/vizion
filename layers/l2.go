package layers

import (
	"bytes"
	"encoding/binary"
	"net"
)

type Ether struct {
	Dst  net.HardwareAddr `field:"Dst"`
	Src  net.HardwareAddr `field:"Src"`
	Type uint16           `field:"Type"`
}

// GetName returns the protocol name.
func (e *Ether) GetName() string {
	return "Ethernet"
}

// Serialize will convert a layer into bytes
func (e *Ether) Build() []byte {

	// Initiate the buffer
	var buffer bytes.Buffer

	// Check for the default value
	if e.Dst == nil {
		e.Dst = net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff}
	}

	// Write data field into the buffer
	binary.Write(&buffer, binary.BigEndian, e.Dst)

	binary.Write(&buffer, binary.BigEndian, e.Src)

	binary.Write(&buffer, binary.BigEndian, e.Type)

	return buffer.Bytes()
}

// Deserialize will convert bytes into a layer
func (e *Ether) Dissect(buf *bytes.Buffer) *bytes.Buffer {
	e.Dst = buf.Next(6)
	e.Src = buf.Next(6)
	e.Type = binary.BigEndian.Uint16(buf.Next(2))

	return buf
}

// BindLayer return the top
func (e *Ether) BindLayer() Layer {

	// If ARP
	if e.Type == 0x0806 {
		return &ARP{}
	}

	return nil
}

var Opcode = map[uint16]string{
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
	Hwtype uint16           `field:"Hwtype"`
	Ptype  uint16           `field:"Ptype"`
	Hwlen  uint8            `field:"Hwlen"`
	Plen   uint8            `field:"Plen"`
	Opcode uint16           `field:"Opcode"`
	Hwsrc  net.HardwareAddr `field:"Hwsrc"`
	Psrc   net.IP           `field:"Psrc"`
	Hwdst  net.HardwareAddr `field:"Hwdst"`
	Pdst   net.IP           `field:"Pdst"`
}

func (a *ARP) GetName() string {
	return "ARP"
}

// TODO: Take into account the addition of paddind depending on frame size
func (a *ARP) Build() []byte {
	// Initiate the buffer
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.BigEndian, a.Hwtype)

	binary.Write(&buffer, binary.BigEndian, a.Ptype)

	binary.Write(&buffer, binary.BigEndian, a.Hwlen)

	binary.Write(&buffer, binary.BigEndian, a.Plen)

	binary.Write(&buffer, binary.BigEndian, a.Opcode)

	binary.Write(&buffer, binary.BigEndian, a.Hwsrc)

	binary.Write(&buffer, binary.BigEndian, a.Psrc)

	binary.Write(&buffer, binary.BigEndian, a.Hwdst)

	binary.Write(&buffer, binary.BigEndian, a.Pdst)

	return buffer.Bytes()
}

func (a *ARP) Dissect(buf *bytes.Buffer) *bytes.Buffer {
	a.Hwtype = binary.BigEndian.Uint16(buf.Next(2))
	a.Ptype = binary.BigEndian.Uint16(buf.Next(2))
	a.Hwlen = uint8(buf.Next(1)[0])
	a.Plen = uint8(buf.Next(1)[0])
	a.Opcode = binary.BigEndian.Uint16(buf.Next(2))
	a.Hwsrc = buf.Next(int(a.Hwlen))
	a.Psrc = buf.Next(int(a.Plen))
	a.Hwdst = buf.Next(int(a.Hwlen))
	a.Pdst = buf.Next(int(a.Plen))
	return buf
}

func (a *ARP) BindLayer() Layer {
	return nil
}

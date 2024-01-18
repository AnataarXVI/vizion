package layers

type TCP struct {
	SrcPort       uint16
	DstPort       uint16
	Seq           uint32
	Ack           uint32
	DataOffset    uint8
	Reserved      uint8
	Flags         uint16
	WindowSize    uint16
	Checksum      uint16
	UrgentPointer uint16
	Options       []byte
	Padding       []byte
}

package layers

import (
	"net"
)

type IP struct {
	Version    uint8  `field:"Version"`
	IHL        uint8  `field:"IHL"`
	TOS        uint8  `field:"TOS"`
	Length     uint16 `field:"Length"`
	ID         uint16 `field:"ID"`
	Flags      uint8  `field:"Flags"`
	FragOffset uint16 `field:"FlagOffset"`
	TTL        uint8  `field:"TTL"`
	Protocol   uint8  `field:"Protocol"`
	Checksum   uint16 `field:"Checksum"`
	Src        net.IP `field:"Src"`
	Dst        net.IP `field:"Dst"`
	Options    []byte `field:"Options"`
	Padding    []byte `field:"Padding"`
}

// GetName retourne le nom du protocole.
func (i *IP) GetName() string {
	return "IPv4"
}

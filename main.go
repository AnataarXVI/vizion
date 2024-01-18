package main

import (
	"fmt"
	"net"

	"github.com/AnataarXVI/vizion/layers"
	"github.com/AnataarXVI/vizion/packet"
)

func main() {

	// Création d'un paquet
	pkt := packet.Packet{}
	dst, err := net.ParseMAC("aa:bb:cc:dd:ee:ff")
	src, err := net.ParseMAC("aa:bb:cc:dd:ee:ff")
	//dst, err := net.ParseMAC("00:11:22:33:44:55")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	// Ajout de la couche Ethernet
	ethernetLayer := &layers.Ether{
		Dst:  dst,
		Src:  src,
		Type: 0x0806,
	}

	pkt.Layers = append(pkt.Layers, ethernetLayer)

	hwsrc, err := net.ParseMAC("aa:bb:cc:dd:ee:ff")
	hwdst, err := net.ParseMAC("aa:bb:cc:dd:ee:ff")
	psrc := net.ParseIP("10.10.10.10")
	pdst := net.ParseIP("10.10.10.10")

	// Ajout de la couche ARP
	arpLayer := &layers.ARP{
		Hwtype: 0x0001,
		Ptype:  0x0800,
		Hwlen:  0x06,
		Plen:   0x04,
		Opcode: 0x0001,
		Hwsrc:  hwsrc,
		Hwdst:  hwdst,
		Psrc:   psrc,
		Pdst:   pdst,
	}

	pkt.Layers = append(pkt.Layers, arpLayer)

	// Modification dynamique du champ SourceIP de la couche IP
	// err = packet.ModifyField(pkt.Layers[0], "Type", uint16(0x0900))
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// Affichage des couches du paquet après la modification
	// pkt.Show()
	// pkt.Build()
	pkt.Show()
	Send(pkt, "wlp1s0")
}

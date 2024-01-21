package main

import (
	"fmt"
	"log"
	"net"
	"syscall"

	"github.com/AnataarXVI/vizion/packet"
	pcap "github.com/packetcap/go-pcap"
)

func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

// Send sends the packet to a raw socket.
func Send(pkt packet.Packet, iface string) error {
	// Build package in bytes
	packetBytes, err := pkt.Build()
	if err != nil {
		return fmt.Errorf("error building packet: %s", err)
	}

	protocol := htons(syscall.ETH_P_IP)
	// Open raw socket
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(protocol))
	if err != nil {
		return fmt.Errorf("error opening raw socket: %s", err)
	}
	defer syscall.Close(fd)

	// Retrieve network interface index
	ifi, err := net.InterfaceByName(iface)
	if err != nil {
		return fmt.Errorf("error getting interface index: %s", err)
	}

	addr := syscall.SockaddrLinklayer{
		Protocol: protocol,
		Ifindex:  ifi.Index,
	}

	// Send bytes to network interface
	err = syscall.Sendto(fd, packetBytes, 0, &addr)
	if err != nil {
		return fmt.Errorf("error sending packet: %s", err)
	}
	fmt.Println("Packet sent successfully.")
	return nil
}

func Sniff(iface string) {

	handle, err := pcap.OpenLive(iface, 1600, true, 0, true)

	if err != nil {
		log.Fatal(err)
	}

	for {
		data, captureInfo, error := handle.ReadPacketData()

		if error != nil {
			log.Fatal(error)
		}

		_ = data
		_ = captureInfo

		// TODO: Insert the captured bytes in the Raw field of a packet and call up the dissection function
		// TODO: Ensure that when Ctrl+C is pressed, the packets received are saved in a revised list at the end.

		// packet.NewPacket(time.Now(), iface, data, data)
	}

}

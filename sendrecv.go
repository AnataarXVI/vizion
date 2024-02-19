package vizion

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/AnataarXVI/vizion/packet"
)

type RawSocket struct {
	protocol uint16
	fd       int
	ifi      *net.Interface
	addr     *syscall.SockaddrLinklayer
}

func htons(i uint16) uint16 {
	return (i<<8)&0xff00 | i>>8
}

// Create and return a Raw socket
func NewSocket(iface string) *RawSocket {

	protocol := htons(syscall.ETH_P_ALL)

	// Open raw socket
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(protocol))
	if err != nil {
		fmt.Printf("error opening raw socket: %s", err)
		return &RawSocket{}
	}

	// Retrieve network interface index
	ifi, err := net.InterfaceByName(iface)
	if err != nil {
		fmt.Printf("error getting interface index: %s", err)
		return &RawSocket{}
	}

	addr := syscall.SockaddrLinklayer{
		Protocol: protocol,
		Ifindex:  ifi.Index,
	}

	return &RawSocket{protocol: protocol, fd: fd, ifi: ifi, addr: &addr}
}

// Send sends the packet to a raw socket.
func Send(pkt packet.Packet, iface string) error {
	// Build package in bytes
	packetBytes, err := pkt.Build()
	if err != nil {
		return fmt.Errorf("error building packet: %s", err)
	}

	// Create the socket
	s := NewSocket(iface)

	defer syscall.Close(s.fd)

	// Send bytes to network interface
	err = syscall.Sendto(s.fd, packetBytes, 0, s.addr)
	if err != nil {
		return fmt.Errorf("error sending packet: %s", err)
	}
	fmt.Println("Packet sent successfully.")
	return nil
}

/*
Sniff will capture traffic on an interface.

The iface parameter corresponding to the interface name.

The prn parameter handle a function called when a packet is received. The function need to handle a packet.Packet struct to get the packet.
*/
func Sniff(iface string, prn func(pkt packet.Packet)) []*packet.Packet {

	var pkt_list []*packet.Packet

	s := NewSocket(iface)
	defer syscall.Close(s.fd)

	// Initiate sigs chan to handle Ctrl+C
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	// Initiate data chan to handle data stream
	data := make(chan packet.Packet)

	// Convert data into Packet and send it in the data chan
	go func() {

		for {
			buffer := make([]byte, 1460)
			n, error := syscall.Read(s.fd, buffer)

			if error != nil {
				log.Fatal(error)
			}

			new_pkt := packet.Packet{Raw: buffer[:n]}
			new_pkt.Dissect()

			// If prn is not nil. It will call the func give in parameters
			if prn != nil {
				prn(new_pkt)
			}

			data <- new_pkt

		}

	}()

	for {

		select {
		// If Interrupt
		case <-sigs:
			return pkt_list
		// Dissect each packet and save it into a list
		case pkt := <-data:
			pkt_list = append(pkt_list, &pkt)
		}

	}

}

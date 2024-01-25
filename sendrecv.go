package main

import (
	"fmt"
	"log"
	"net"
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

	protocol := htons(syscall.ETH_P_IP)

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

func Sniff(iface string) {

	s := NewSocket(iface)
	defer syscall.Close(s.fd)

	// pkt_list := make(chan []packet.Packet)

	// c := make(chan os.Signal)
	// signal.Notify(c, syscall.SIGINT)

	// the original handling for that signal will be reinstalled, restoring the non-Go signal handler if any.
	// defer signal.Reset(syscall.SIGINT)

	//go func() {

	for {
		buffer := make([]byte, 1460)

		n, error := syscall.Read(s.fd, buffer)

		if error != nil {
			log.Fatal(error)
		}

		// TODO: Insert the captured bytes in the Raw field of a packet and call up the dissection function
		// TODO: Ensure that when Ctrl+C is pressed, the packets received are saved in a revised list at the end.

		fmt.Println(buffer[:n])

		//pkt_list <- packet.Packet{Raw: buffer[:n]}

	}

	// }()

	// wait go routine to finish or signal received
	// <-c
	// list := <-pkt_list
	// return list
}

/*
Objectifs:
  - gestion des sockets
  - capture le trafic
  - injecte du trafic

Fonctions:
  - Sniff : Lance la capture le trafic
  - Recv 	: Reçoit les données
  - Send	: Envoie des données
  - InitHandler : Revoie un handle

Structure:
  - Handle
*/

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

// Send envoie le paquet sur une raw socket.
func Send(pkt packet.Packet, iface string) error {
	// Build le paquet en bytes
	packetBytes, err := pkt.Build()
	if err != nil {
		return fmt.Errorf("error building packet: %s", err)
	}

	protocol := htons(syscall.ETH_P_IP)
	// Ouvrir la raw socket
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, int(protocol))
	if err != nil {
		return fmt.Errorf("error opening raw socket: %s", err)
	}
	defer syscall.Close(fd)

	// Récupérer l'index de l'interface réseau
	ifi, err := net.InterfaceByName(iface)
	if err != nil {
		return fmt.Errorf("error getting interface index: %s", err)
	}
	// Configurer les options de la raw socket
	// err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_BINDTODEVICE, ifi.Index)
	// if err != nil {
	// 	return fmt.Errorf("error setting socket option: %s", err)
	// }

	addr := syscall.SockaddrLinklayer{
		Protocol: protocol,
		Ifindex:  ifi.Index,
	}

	// Envoyer les bytes sur l'interface réseau
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

		// TODO: Insérer les octets capturés dans le champ Raw d'un paquet et faire appel à la fonction de dissection
		// TODO: Faire en sorte qu'en cas de Ctrl+C les paquets reçus soient enregistrés dans une liste revoyé à la fin

		// packet.NewPacket(time.Now(), iface, data, data)
	}

}

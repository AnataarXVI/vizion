package layers

import (
	"math/rand"
	"net"

	"github.com/AnataarXVI/vizion/buffer"
)

type IP struct {
	Version    uint8
	IHL        uint8
	TOS        uint8
	Length     uint16
	ID         uint16
	Flags      uint8
	FragOffset uint8
	TTL        uint8
	Protocol   uint8
	Checksum   uint16
	Src        net.IP
	Dst        net.IP
}

type IP_Options struct {
	Type   uint8
	Length uint8
	Data   uint8
}

func (i *IP) GetName() string {
	return "IPv4"
}

func IPv4Layer() IP {
	ip := IP{}
	ip.SetDefault()
	return ip
}

// IP Options
var IPOPTIONS = map[uint8]string{
	0:  "end_of_list",
	1:  "nop",
	2:  "security",
	3:  "loose_source_route",
	4:  "timestamp",
	5:  "extended_security",
	6:  "commercial_security",
	7:  "record_route",
	8:  "stream_id",
	9:  "strict_source_route",
	10: "experimental_measurement",
	11: "mtu_probe",
	12: "mtu_reply",
	13: "flow_control",
	14: "access_control",
	15: "encode",
	16: "imi_traffic_descriptor",
	17: "extended_IP",
	18: "traceroute",
	19: "address_extension",
	20: "router_alert",
	21: "selective_directed_broadcast_mode",
	23: "dynamic_packet_state",
	24: "upstream_multicast_packet",
	25: "quick_start",
	30: "rfc4727_experiment",
}

// IHL Lengths
var IHLLENGTH = map[uint8]string{
	5:  "20 bytes",
	6:  "24 bytes",
	7:  "28 bytes",
	8:  "32 bytes",
	9:  "36 bytes",
	10: "40 bytes",
	11: "44 bytes",
	12: "48 bytes",
	13: "52 bytes",
	14: "56 bytes",
	15: "60 bytes",
}

func (i *IP) SetDefault() {

	netAddr, _ := net.InterfaceAddrs()

	i.Version = 4
	i.IHL = 5
	i.TOS = 20
	i.Length = 20
	i.ID = uint16(rand.Intn(65535))
	i.Flags = 2
	i.FragOffset = 0
	i.TTL = 64
	i.Protocol = 6
	i.Checksum = uint16(rand.Intn(65535))
	i.Dst = net.IP{0, 0, 0, 0}
	i.Src = netAddr[1].(*net.IPNet).IP.To4()
}

func (i *IP) Build() *buffer.ProtoBuff {
	// Initiate the buffer
	var buffer buffer.ProtoBuff

	buffer.AddBitsUint8([]string{"Version", "IHL"}, []uint8{i.Version, i.IHL}, []int{4, 4}, []any{nil, IHLLENGTH[i.IHL]})
	buffer.Add("TOS", i.TOS, nil)
	buffer.Add("Length", i.Length, nil)
	buffer.Add("ID", i.ID, nil)
	buffer.Add("Flags", i.Flags, nil)
	buffer.Add("FragOffset", i.FragOffset, nil)
	buffer.Add("TTL", i.TTL, nil)
	buffer.Add("Protocol", i.Protocol, nil)
	buffer.Add("Checksum", i.Checksum, nil)
	buffer.Add("Dst", i.Dst, nil)
	buffer.Add("Src", i.Src, nil)

	return &buffer
}

func (i *IP) Dissect(buffer *buffer.ProtoBuff) *buffer.ProtoBuff {

	first_byte := buffer.NextUint8()

	i.Version = (first_byte & 0xF0) >> 4
	i.IHL = first_byte & 0x0F
	i.TOS = buffer.NextUint8()
	i.Length = buffer.NextUint16()
	i.ID = buffer.NextUint16()
	i.Flags = buffer.NextUint8()
	i.FragOffset = buffer.NextUint8()
	i.TTL = buffer.NextUint8()
	i.Protocol = buffer.NextUint8()
	i.Checksum = buffer.NextUint16()
	i.Dst = buffer.Next(4)
	i.Src = buffer.Next(4)

	return buffer
}

func (i *IP) BindLayer() Layer {
	return nil
}

// Create and return an IP layer with default value set
func IPLayer() IP {
	ip := IP{}
	ip.SetDefault()
	return ip
}

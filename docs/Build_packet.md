# How to build a packet ?

## Quick setup

To use the `vizion` library, you need to import it into your code as follows: 

```go
import (
    "github.com/AnataarXVI/vizion"
    "github.com/AnataarXVI/vizion/packet"
    "github.com/AnataarXVI/vizion/layers"
)
```

Next, run the command `go mod tidy` to install the library. Dependencies will be installed automatically.


## First step

First, you need to create a packet using the `Packet` structure.

```go
pkt := packet.Packet{}
```

For the moment, the packet is empty and contains no layers. 

## Second step

Now you need to create the layers you want to add to your packet. For this example, we'll create an ARP packet. To do this, we need to create an Ethernet layer and an ARP layer.

_For the Ethernet Layer :_

```go
dst, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")
src, _ := net.ParseMAC("11:22:33:44:55:66")

// Ethernet Layer
ethernetLayer := layers.Ether{
    Dst:  dst,
    Src:  src,
    Type: 0x0806,
}
```

_For the ARP Layer :_

```go

hwsrc, _ := net.ParseMAC("aa:bb:cc:dd:ee:ff")
hwdst, _ := net.ParseMAC("11:22:33:44:55:66")
psrc := net.ParseIP("10.10.10.10")
pdst := net.ParseIP("10.10.10.20")

// ARP Layer
arpLayer := layers.ARP{
	Hwtype: 0x0001,
	Ptype:  0x0800,
	Hwlen:  6,
	Plen:   4,
	Opcode: 0x0001,
	Hwsrc:  hwsrc,
	Hwdst:  hwdst,
	Psrc:   psrc,
	Pdst:   pdst,
}
```

## Third step

All we have to do now is add our layers to our packet !


```go
pkt.AddLayers(&ethernetLayer, &arpLayer)
```

## Display a packet

You can display your packet to view its composition. 

```go
pkt.Show()
```

```
###[ Ethernet ]###
	Dst = aa:bb:cc:dd:ee:ff
	Src = 11:22:33:44:55:66
	Type = 2054
###[ ARP ]###
	Hwtype = 1
	Ptype = 2048
	Hwlen = 6
	Plen = 4
	Opcode = 1
	Hwsrc = aa:bb:cc:dd:ee:ff
	Psrc = 10.10.10.10
	Hwdst = 11:22:33:44:55:66
	Pdst = 10.10.10.20
```

## Sending packet

The `Send()` function is used to send a packet, indicating the interface.

```go
vizion.Send(pkt, "lo")
```

```
Packet sent successfully.
```
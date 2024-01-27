# Vizion
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](LICENSE)

This tool is designed for offensive use. It enables packets to be forged and quickly dissected. In the future, it will incorporate attack functions such as mitm and others.

## Getting started

Vizion can be imported as a library. This library lets you create fully customizable packets. 

Here's how to import the library:

```go
import (
    "github.com/AnataarXVI/vizion"
    "github.com/AnataarXVI/vizion/packet"
    "github.com/AnataarXVI/vizion/layers"
)
```


Here's an example of how to create a package: 

```go
// Initialize the packet
pkt := packet.Packet{}
```

You can easily add layers to the packet:

```go
dst, err := net.ParseMAC("aa:bb:cc:dd:ee:ff")
src, err := net.ParseMAC("aa:bb:cc:dd:ee:ff")

// First we should create the layer
ethernetLayer := layers.Ether{
    Dst:  dst,
    Src:  src,
    Type: 0x0806,
}

// Then we add the layer to the packet
pkt.AddLayers(&ethernetLayer)
```

To see the packet structure you can use the `Show()` function on the packet.

```go
pkt.Show()

>>>
###[ Ethernet ]###
	Dst = aa:bb:cc:dd:ee:ff
	Src = aa:bb:cc:dd:ee:ff
	Type = 2054
```

Finally, to send the packet, we use the `Send()` function.

```go
Send(pkt,"lo")

>>>
Packet sent successfully.
```

## Ressources

In the [docs](./docs) folder, you'll find all the documentation you need to get started with the library.

Some useful resources :

- [Build a packet](./docs/Build_packet.md)
- [Sniff packets](./docs/Sniff_packet.md)
- [Create a dissector](./docs/Create_dissector.md)



## Contributing

This project is currently under development. I am open to any improvements or advice. 

You can contact me at anataar@protonmail.com or open an issue.
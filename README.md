# <img src="docs/vizion.png" width="100" valign="middle" alt="Vizion" /> &nbsp;&nbsp; Vizion

[![GoDoc](https://godoc.org/github.com/google/gopacket?status.svg)](https://godoc.org/github.com/AnataarXVI/vizion)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](LICENSE)

Vizion is a library designed for building and dissecting network packages with great speed and efficiency.

## Installation 

To download this library do :

```
$ go get github.com/AnataarXVI/vizion
```

## Getting started

Vizion can be imported as a library. This library lets you create fully customizable packets. 

Here's how to import the library:

```go
import (
    . "github.com/AnataarXVI/vizion"
    "github.com/AnataarXVI/vizion/packet"
    "github.com/AnataarXVI/vizion/layers"
)
```

Here is an example of how to create a packet :

```go
// Create the packet
pkt := packet.Packet{}

// Create the Ethernet layer with default values set
etherLayer := layers.EtherLayer()
// Modify the Type field of the layer
etherLayer.Type = 0x0806

// Create the ARP layer with default values set 
arpLayer := layers.ARPLayer()

// Add layers to the packet
pkt.AddLayers(&etherLayer, &arpLayer)

// Show the packet composition
pkt.Show()

// Send the packet on 'lo' interface
Send(pkt, "lo")
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
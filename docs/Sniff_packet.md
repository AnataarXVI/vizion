# How to sniff packet ?

## Quick setup

To use the `vizion` library, you need to import it into your code as follows: 

```go
import (
    "github.com/AnataarXVI/vizion"
)
```

Next, run the command `go mod tidy` to install the library. Dependencies will be installed automatically.

## Sniff

The `Sniff()` function captures traffic on a given interface.

```go
func main() {
    sniffed_pkts := vizion.Sniff("lo", nil)
}
```

The function takes as argument the name of the interface and optionally a function that will be called each time a packet is received. The default value is nil.

When the capture is interrupted, all packets are returned by the function and can be further manipulated.

Here's a simple example of how to use a function on the fly:

```go
package main

import (
    "github.com/AnataarXVI/vizion"
    "github.com/AnataarXVI/vizion/packet"
)

func Display(pkt packet.Packet) {
    pkt.Show()
}

func main() {
    sniffed_pkts := vizion.Sniff("lo", Display)
}
```

**Be careful !** The intermediate function must take a `packet.Packet` type as parameter.
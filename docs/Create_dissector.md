# How to create a dissector ?

This library uses dissectors. Each dissector corresponds to a protocol. They are used to build and dissect packages. They are located in the [layers](../layers/) folder.

Each layer is represented by a structure with several fields corresponding to the protocol. For each field, you must indicate which type it is.  

Let's take the ARP layer as an example:

```go
// ARP Layer
type ARP struct {
	Hwtype uint16           `field:"Hwtype"`
	Ptype  uint16           `field:"Ptype"`
	Hwlen  uint8            `field:"Hwlen"`
	Plen   uint8            `field:"Plen"`
	Opcode uint16           `field:"Opcode"`
	Hwsrc  net.HardwareAddr `field:"Hwsrc"`
	Psrc   net.IP           `field:"Psrc"`
	Hwdst  net.HardwareAddr `field:"Hwdst"`
	Pdst   net.IP           `field:"Pdst"`
}
```

As you can see, the ARP layer is a structure with different fields. 

For the structure to be displayed correctly when using the `Show()` function on the packet, it is necessary to assign a tag to the name of `field` with the field name as its value. 

**This name must be the same as that used in the structure.**

The following functions must be defined for each layer: `GetName()`, `SetDefault()` , `Build()`, `Dissect()`, `BindLayer()` and `<LayerName>Layer()`.

## GetName

This function is used to specify the layer name. It is used in the `Show()` function to retrieve the name of the current layer. The function returns a string of characters.

```go
func (a *ARP) GetName() string {
	return "ARP"
}
```

## SetDefault

The purpose of this function is to define default values for each protocol field.

```go
func (a *ARP) SetDefault() {

	ifaces, _ := net.Interfaces()
	netAddr, _ := net.InterfaceAddrs()

	a.Hwtype = 1
	a.Ptype = 0x0800
	a.Hwlen = 6
	a.Plen = 4
	a.Opcode = 1
	a.Hwsrc = ifaces[1].HardwareAddr
	a.Psrc = netAddr[1].(*net.IPNet).IP.To4()
	a.Hwdst = net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	a.Pdst = net.IPv4(0, 0, 0, 0)
}
```


## Build

This function is used when creating a packet using the ARP layer. It returns a corresponding array of bytes containing the values of each field.

```go
func (a *ARP) Build() []byte {
    // Process
}
```

The first step is to initialize a buffer containing the bytes of our fields.

```go
// Initiate the buffer
var buffer bytes.Buffer
```

Next, we need to add the value of each of our fields inside the buffer. 

```go
// Add Hwtype field into the buffer
binary.Write(&buffer, binary.BigEndian, a.Hwtype)
// Add Ptype field into the buffer
binary.Write(&buffer, binary.BigEndian, a.Ptype)
// Add Hwlen field into the buffer
binary.Write(&buffer, binary.BigEndian, a.Hwlen)
// Add Plen field into the buffer
binary.Write(&buffer, binary.BigEndian, a.Plen)
// Add Opcode field into the buffer
binary.Write(&buffer, binary.BigEndian, a.Opcode)
// Add Hwsrc field into the buffer
binary.Write(&buffer, binary.BigEndian, a.Hwsrc)
// Add Psrc field into the buffer
binary.Write(&buffer, binary.BigEndian, a.Psrc)
// Add Hwdst field into the buffer
binary.Write(&buffer, binary.BigEndian, a.Hwdst)
// Add Pdst field into the buffer
binary.Write(&buffer, binary.BigEndian, a.Pdst)
```

Finally, we return the contents of our buffer.

```go
return buffer.Bytes()
```

## Dissect

This function converts an array of bytes into a layer. It takes as argument a buffer corresponding to the undissected bytes and returns the buffer.

```go
func (a *ARP) Dissect(buf *bytes.Buffer) *bytes.Buffer {
    // Process
    return buf
}
```

Bytes are inserted for each field in the layer. 
**Be careful with the type !**

```go
// Inserts bytes in Hwtype
a.Hwtype = binary.BigEndian.Uint16(buf.Next(2))
// Inserts bytes in Ptype
a.Ptype = binary.BigEndian.Uint16(buf.Next(2))
// Inserts byte in Hwlen
a.Hwlen = uint8(buf.Next(1)[0])
// Inserts byte in Plen
a.Plen = uint8(buf.Next(1)[0])
// Inserts bytes in Opcode
a.Opcode = binary.BigEndian.Uint16(buf.Next(2))
// Inserts bytes in Hwsrc
a.Hwsrc = buf.Next(int(a.Hwlen))
// Inserts bytes in Psrc
a.Psrc = buf.Next(int(a.Plen))
// Inserts bytes in Hwdst
a.Hwdst = buf.Next(int(a.Hwlen))
// Inserts bytes in Pdst
a.Pdst = buf.Next(int(a.Plen))
```

## BindLayer

This function is used to indicate the next layer. It returns a Layer strucuture corresponding to the next layer.

```go
func (a *ARP) BindLayer() Layer {
	return nil
}
```

If there is no next layer, the value `nil` is returned. 

Let's imagine that the next layer is IP (this is an example, of course ^^), then we'd have :

```go
// Condition on the field indicating the next layer
if a.field == value {
    return &IP{}
}
```

Once you've understood this, it's important to specify the condition for indicating the presence of our layer in the `BindLayer()` function of the previous layer !

## <LayerName>Layer

This is the function that will be called when the layer is created. It instantiates the layer and applies default values to the fields by calling the `SetDefault()` function. Finally, it returns the layer, which can then be manipulated by the user.

```go
func ARPLayer() ARP {
	arp := ARP{}
	arp.SetDefault()
	return arp
}
```
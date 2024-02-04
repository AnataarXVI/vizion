package layers

import (
	"bytes"
	"encoding/binary"
)

type Raw struct {
	Load []byte `field:"Load"`
}

func (r *Raw) GetName() string {
	return "Raw"
}

func RawLayer() Raw {
	return Raw{}
}

func (r *Raw) SetDefault() {

}

func (r *Raw) Build() []byte {
	// Initiate the buffer
	var buffer bytes.Buffer

	binary.Write(&buffer, binary.BigEndian, r.Load)

	return buffer.Bytes()
}

func (r *Raw) Dissect(buf *bytes.Buffer) *bytes.Buffer {

	r.Load = buf.Bytes()

	return nil
}

func (r *Raw) BindLayer() Layer {
	return nil
}

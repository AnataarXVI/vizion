package layers

import (
	"github.com/AnataarXVI/vizion/buffer"
)

type Raw struct {
	Load []byte
}

func (r *Raw) GetName() string {
	return "Raw"
}

func RawLayer() Raw {
	return Raw{}
}

func (r *Raw) SetDefault() {

}

func (r *Raw) Build() *buffer.ProtoBuff {
	// Initiate the buffer
	var buffer buffer.ProtoBuff

	buffer.Add("Raw", r.Load, nil)

	return &buffer
}

func (r *Raw) Dissect(buf *buffer.ProtoBuff) *buffer.ProtoBuff {

	r.Load = buf.Bytes()

	return nil
}

func (r *Raw) BindLayer() Layer {
	return nil
}

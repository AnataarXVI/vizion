package layers

import (
	"github.com/AnataarXVI/vizion/buffer"
)

// Layer is an interface representing a layer of the package.
type Layer interface {
	GetName() string
	SetDefault()
	Build() *buffer.ProtoBuff
	Dissect(*buffer.ProtoBuff) *buffer.ProtoBuff
	BindLayer() Layer
}

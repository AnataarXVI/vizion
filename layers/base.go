package layers

import (
	"bytes"
)

// Layer is an interface representing a layer of the package.
type Layer interface {
	GetName() string
	SetDefault()
	Build() []byte
	Dissect(*bytes.Buffer) *bytes.Buffer
	BindLayer() Layer
}

package layers

import (
	"bytes"
)

// Layer is an interface representing a layer of the package.
type Layer interface {
	GetName() string
	Build() []byte
	Dissect(*bytes.Buffer) *bytes.Buffer
	BindLayer() Layer
}

// TODO: Find a way to set default fields during package build

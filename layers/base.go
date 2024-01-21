package layers

import (
	"bytes"
)

// Layer is an interface representing a layer of the package.
type Layer interface {
	GetName() string
	Build() []byte

	// TODO: Add the layer to the list of layers in the package
	Dissect(*bytes.Buffer) *bytes.Buffer
}

// TODO: Add a bind_layer function to link two layers together

// TODO: Find a way to set default fields during package build

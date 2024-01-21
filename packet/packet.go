package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/AnataarXVI/vizion/layers"
)

// Packet is a structure representing a network packet.
type Packet struct {
	Layers []layers.Layer
	Raw    []byte
}

// DisplayLayers displays the types of all layers in the package and accesses specific fields via the reflect lib.
func (p *Packet) Show() {
	for _, layer := range p.Layers {
		// Display the actual layer
		fmt.Printf("###[ %s ]###\n", layer.GetName())
		layerType := reflect.TypeOf(layer).Elem()
		layerValue := reflect.ValueOf(layer).Elem()

		for i := 0; i < layerType.NumField(); i++ {
			fieldName := layerType.Field(i).Tag.Get("field")
			fieldValue := layerValue.Field(i).Interface()
			fmt.Printf("\t%s = %v\n", fieldName, fieldValue)

		}
	}
}

// ModifyField dynamically modifies a layer field.
func ModifyField(layer layers.Layer, fieldName string, value interface{}) error {
	layerValue := reflect.ValueOf(layer).Elem()
	field := layerValue.FieldByName(fieldName)
	if !field.IsValid() {
		return fmt.Errorf("unknown field: %s", fieldName)
	}

	// Vérifier le type de la valeur
	if field.Type() != reflect.TypeOf(value) {
		return fmt.Errorf("invalid type for field %s, expected %s, got %s", fieldName, field.Type(), reflect.TypeOf(value))
	}

	// Check value type
	field.Set(reflect.ValueOf(value))

	return nil
}

// Dissect converts received bytes into Layer
func (p *Packet) Dissect() {
	var buffer bytes.Buffer

	// FIXME: Not good, because the layers won't be defined beforehand. Hence the use of bind_layer
	for _, layer := range p.Layers {
		bytes_remaining := layer.Dissect(&buffer)
		_ = bytes_remaining
	}
}

// Build converts the packet into a sequence of bytes.
func (p *Packet) Build() ([]byte, error) {
	var buffer bytes.Buffer
	for _, layer := range p.Layers {
		serialized_data := layer.Build()

		err := binary.Write(&buffer, binary.BigEndian, serialized_data)
		if err != nil {
			return nil, fmt.Errorf("error writing layer field: %s", err)
		}

	}
	return buffer.Bytes(), nil
}

func (p *Packet) Bind_layer() {

}

func (p *Packet) AddLayers(layers ...layers.Layer) {
	for _, layer := range layers {
		p.Layers = append(p.Layers, layer)
	}
}

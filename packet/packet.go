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
	fmt.Print("\n")
}

// ModifyField dynamically modifies a layer field.
func ModifyField(layer layers.Layer, fieldName string, value interface{}) error {

	// TODO: Modify Raw field of the packet with the new value

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

	// Insert Raw data into the buffer
	binary.Write(&buffer, binary.BigEndian, p.Raw)

	// While all bytes aren't convert into layers
	for {

		// If buffer is nil
		if buffer.Len() == 0 {
			break
		}

		// Start with the Ethernet layer
		if len(p.Layers) == 0 {

			// Add Ethernet layer into p.Layers
			p.Layers = append(p.Layers, &layers.Ether{})

			// Do the dissection
			bytes_remaining := p.Layers[0].Dissect(&buffer)

			// Call the BindLayer func to know the next layer
			next_layer := p.Layers[0].BindLayer()

			// Add the next Layer into p.Layers
			if next_layer != nil && buffer.Len() != 0 {
				p.Layers = append(p.Layers, next_layer)

				// Return Raw Layer if the buffer isn't nil and the next layer is nil
			} else if next_layer == nil && buffer.Len() != 0 {
				p.Layers = append(p.Layers, &layers.Raw{Load: bytes_remaining.Bytes()})
				break

				// There is no next_layer and all bytes are dissect
			} else {
				break
			}

			// Remove used bytes
			buffer = *bytes_remaining

		} else { // Other Layers

			// Get the last layer of p.Layers
			// Do the dissection
			bytes_remaining := p.Layers[len(p.Layers)-1].Dissect(&buffer)

			// Call the BindLayer func to know the next layer
			next_layer := p.Layers[len(p.Layers)-1].BindLayer()

			if next_layer != nil && len(buffer.Bytes()) != 0 {
				p.Layers = append(p.Layers, next_layer)

				// Return Raw Layer if the buffer isn't nil and the next layer is nil
			} else if next_layer == nil && buffer.Len() != 0 {
				p.Layers = append(p.Layers, &layers.Raw{Load: bytes_remaining.Bytes()})
				break

				// There is no next_layer and all bytes are dissect
			} else {
				break
			}

			// Add the next Layer into p.Layers
			p.Layers = append(p.Layers, next_layer)
			// Remove used bytes
			buffer = *bytes_remaining
		}
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

func (p *Packet) AddLayers(layers ...layers.Layer) {
	p.Layers = append(p.Layers, layers...)
}

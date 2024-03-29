package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/AnataarXVI/vizion/buffer"
	"github.com/AnataarXVI/vizion/layers"
	"github.com/AnataarXVI/vizion/utils"
)

// Packet is a structure representing a network packet.
type Packet struct {
	Layers []layers.Layer
	Raw    []byte
}

// Show displays the types of all layers in the package and accesses specific fields via the reflect lib.
func (p *Packet) Show() {

	for _, layer := range p.Layers {
		// Display the actual layer
		utils.Display_Layer(layer.GetName())

		ProtoBuff := layer.Build()

		loaded_fields := ProtoBuff.GetLoadedFields()

		var actual_sublayer string
		var cache string

		for _, field := range loaded_fields {

			// Check if the field belongs to a Sublayer
			if !reflect.DeepEqual(field.ParentLayer, buffer.LoadedSubLayer{}) {

				// Store the Sublayer
				actual_sublayer = field.ParentLayer.LayerName

				// Check if the previous field haven't the same Sublayer
				if actual_sublayer != cache {
					utils.Display_SubLayer(field.ParentLayer.Name)
					utils.Display_SubFields(field.Name, field.Value, field.Enum)
				} else { // Previous field are in the same Sublayer
					utils.Display_SubFields(field.Name, field.Value, field.Enum)
				}

				// store the actual Sublayer into a cache
				cache = actual_sublayer

				//fmt.Println(actual_sublayer, cache)

			} else { // Simple field
				utils.Display_Fields(field.Name, field.Value, field.Enum)
			}

		}

	}
	fmt.Print("\n")
}

// ShowF displays only layers passed as arguments. Other layers are not detailed.
func (p *Packet) ShowF(filter ...string) {

	for _, layer := range p.Layers {
		utils.Display_Layer(layer.GetName())
		for _, fl := range filter {
			if fl == layer.GetName() {
				ProtoBuff := layer.Build()
				loaded_fields := ProtoBuff.GetLoadedFields()

				for _, field := range loaded_fields {
					if !reflect.DeepEqual(field.ParentLayer, buffer.LoadedSubLayer{}) {
						utils.Display_SubFields(field.Name, field.Value, field.Enum)
					} else {
						utils.Display_Fields(field.Name, field.Value, field.Enum)
					}
				}
			}
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
	var buffer buffer.ProtoBuff

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

func (p *Packet) DissectFrom() {

}

// Build converts the packet into a sequence of bytes.
func (p *Packet) Build() ([]byte, error) {
	var buffer bytes.Buffer
	for _, layer := range p.Layers {
		serialized_data := layer.Build()

		err := binary.Write(&buffer, binary.BigEndian, serialized_data.Bytes())
		if err != nil {
			return nil, fmt.Errorf("error writing layer field: %s", err)
		}

	}
	return buffer.Bytes(), nil
}

func (p *Packet) AddLayers(layers ...layers.Layer) {
	p.Layers = append(p.Layers, layers...)
}

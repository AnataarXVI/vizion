/*
Objectifs:
  - Traitement des paquets
  - Affichage du paquet
  - Operation sur le paquet

Fonctions:
  - Show() / Show2()
  - Do_Build()
  - Do_Dissect()
  - Pre_Build
  - Post_Build
  - Etc..


Structure:
  - Packet
*/

package packet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/AnataarXVI/vizion/layers"
)

// Packet est une structure représentant un paquet réseau.
type Packet struct {
	Layers []layers.Layer
	Raw    []byte
}

// DisplayLayers affiche les types de toutes les couches du paquet et accède aux champs spécifiques via la réflexion.
func (p *Packet) Show() {
	for _, layer := range p.Layers {
		// Affiche la couche actuelle
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

// ModifyField permet de modifier dynamiquement un champ de la couche.
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

	// Mettre à jour la valeur du champ
	field.Set(reflect.ValueOf(value))

	return nil
}

// Dissect convertit les octets reçus en Layer
func (p *Packet) Dissect() {
	var buffer bytes.Buffer

	// FIXME: Pas bon car les layers ne seront pas définies au préalable. D'où l'utilisation du bind_layer
	for _, layer := range p.Layers {
		bytes_remaining := layer.Dissect(&buffer)
		_ = bytes_remaining
	}
}

// Build convertit le paquet en une séquence de bytes.
func (p *Packet) Build() ([]byte, error) {
	var buffer bytes.Buffer
	for _, layer := range p.Layers {
		serialized_data := layer.Build()

		err := binary.Write(&buffer, binary.BigEndian, serialized_data)
		if err != nil {
			return nil, fmt.Errorf("error writing layer field: %s", err)
		}

		// layerType := reflect.TypeOf(layer).Elem()
		// layerValue := reflect.ValueOf(layer).Elem()

		// for i := 0; i < layerType.NumField(); i++ {
		// 	value := layerValue.Field(i).Interface()
		// 	err := binary.Write(&buffer, binary.BigEndian, value)
		// 	if err != nil {
		// 		return nil, fmt.Errorf("error writing layer field: %s", err)
		// 	}
		// }
	}
	return buffer.Bytes(), nil
}

func (p *Packet) Bind_layer() {

}

func (p *Packet) AddLayers(layers ...layers.Layer) {
	for i := len(layers) - 1; i >= 0; i-- {
		p.Layers = append(p.Layers, layers[i])
	}
}

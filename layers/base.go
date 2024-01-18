package layers

import (
	"bytes"
)

// Layer est une interface représentant une couche du paquet.
type Layer interface {
	GetName() string
	Build() []byte

	// TODO: Ajouter la layer dans la liste layers du paquet
	Dissect(*bytes.Buffer) *bytes.Buffer
}

// TODO: Ajouter une fonction bind_layer permettant de relier deux couches entres elles

// TODO: Trouver un moyen de mettre des champs par défaut lors du build de paquet

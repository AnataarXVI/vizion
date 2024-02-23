package buffer

import (
	"bytes"
	"encoding/binary"
)

type ProtoBuff struct {
	bytes.Buffer
	loaded_fields []LoadedField // List contians loaded fields after building packet
}

func (buffer *ProtoBuff) GetLoadedFields() []LoadedField {
	return buffer.loaded_fields
}

// Add will write into the buffer the value passed in args.
// It will also save the field for the Show function.
func (buffer *ProtoBuff) Add(fieldname string, value any, enum any) {
	binary.Write(buffer, binary.BigEndian, value)
	// If enum and key exist
	if enum != nil && enum != "" {
		buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value, Enum: enum})
	} else {
		buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value})
	}

}

func (buffer *ProtoBuff) Add_Le(fieldname string, value any, enum any) {
	binary.Write(buffer, binary.LittleEndian, value)
	// If enum and key exist
	if enum != nil && enum != "" {
		buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value, Enum: enum})
	} else {
		buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value})
	}

}

type LoadedField struct {
	Name  string
	Value any
	Enum  any
}

func (f *LoadedField) GetName() string {
	return f.Name
}

func (f *LoadedField) GetValue() any {
	return f.Value
}

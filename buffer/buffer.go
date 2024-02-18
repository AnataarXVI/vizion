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

func (buffer *ProtoBuff) Add(fieldname string, value any) {
	binary.Write(buffer, binary.BigEndian, value)
	buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value})
}

// func (b *ProtoBuff) Add_Le(field any) {
// 	binary.Write(b, binary.LittleEndian, field)
// }

type LoadedField struct {
	Name  string
	Value any
}

func (f *LoadedField) GetName() string {
	return f.Name
}

func (f *LoadedField) GetValue() any {
	return f.Value
}

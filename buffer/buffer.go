package buffer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/bits"
	"reflect"
)

type ProtoBuff struct {
	bytes.Buffer
	loaded_fields []LoadedField // List contains loaded fields after building packet
}

func (buffer *ProtoBuff) GetLoadedFields() []LoadedField {
	return buffer.loaded_fields
}

// Add will write into the buffer the value passed in args.
// It will also save the field for the Show function.
func (buffer *ProtoBuff) Add(fieldname string, value any, enum any) {

	if reflect.TypeOf(value).Kind() == reflect.String { // If value is string
		// Add field value to buffer
		buffer.WriteString(value.(string))

		// If enum and key exist
		if enum != nil && enum != "" {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value, Enum: enum})
		} else {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value})
		}

	} else { // For other data types

		// Add field value to buffer
		binary.Write(buffer, binary.BigEndian, value)

		// If enum and key exist
		if enum != nil && enum != "" {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value, Enum: enum})
		} else {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value})
		}
	}

	//fmt.Println(buffer.loaded_fields)

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

func (buffer *ProtoBuff) AddBitsUint8(fieldname []string, values []uint8, sizes []int, enum []any) {

	if len(values) != len(sizes) {
		fmt.Errorf("Error: values and sizes must be same size")
	}

	var result uint8
	var offset int
	for i := 0; i < len(values); i++ {
		if i == 0 {
			result = (bits.RotateLeft8(values[i], -sizes[i]))
			offset += sizes[i]
		} else {
			result |= bits.RotateLeft8(values[i], -(offset + sizes[i]))
			offset += sizes[i]
		}

		if enum != nil && enum[i] != nil && enum[i] != "" {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname[i], Value: values[i], Enum: enum[i]})
		} else {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname[i], Value: values[i]})
		}

	}

	binary.Write(buffer, binary.BigEndian, result)
}

func (buffer *ProtoBuff) AddBitsUint16(fieldname []string, values []uint16, sizes []int, enum []any) {

	if len(values) != len(sizes) {
		fmt.Errorf("Error: values and sizes must be same size")
	}

	var result uint16
	var offset int
	for i := 0; i < len(values); i++ {

		if i == 0 {
			result = (bits.RotateLeft16(values[i], -sizes[i]))
			offset += sizes[i]
		} else {
			result |= bits.RotateLeft16(values[i], -(offset + sizes[i]))
			offset += sizes[i]
		}

		if enum != nil && enum[i] != nil && enum[i] != "" {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname[i], Value: values[i], Enum: enum[i]})
		} else {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname[i], Value: values[i]})
		}

	}
	binary.Write(buffer, binary.BigEndian, result)
}

func (buffer *ProtoBuff) AddBitsUint32(fieldname []string, values []uint32, sizes []int, enum []any) {

	if len(values) != len(sizes) {
		fmt.Errorf("Error: values and sizes must be same size")
	}

	var result uint32
	var offset int
	for i := 0; i < len(values); i++ {

		if i == 0 {
			result = (bits.RotateLeft32(values[i], -sizes[i]))
			offset += sizes[i]
		} else {
			result |= bits.RotateLeft32(values[i], -(offset + sizes[i]))
			offset += sizes[i]
		}

		if enum != nil && enum[i] != nil && enum[i] != "" {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname[i], Value: values[i], Enum: enum[i]})
		} else {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname[i], Value: values[i]})
		}

	}
	binary.Write(buffer, binary.BigEndian, result)
}

func (buffer *ProtoBuff) AddBitsUint64(fieldname []string, values []uint64, sizes []int, enum []any) {

	if len(values) != len(sizes) {
		fmt.Errorf("Error: values and sizes must be same size")
	}

	var result uint64
	var offset int
	for i := 0; i < len(values); i++ {

		if i == 0 {
			result = (bits.RotateLeft64(values[i], -sizes[i]))
			offset += sizes[i]
		} else {
			result |= bits.RotateLeft64(values[i], -(offset + sizes[i]))
			offset += sizes[i]
		}

		if enum != nil && enum[i] != nil && enum[i] != "" {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname[i], Value: values[i], Enum: enum[i]})
		} else {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname[i], Value: values[i]})
		}

	}
	binary.Write(buffer, binary.BigEndian, result)
}

func (buffer *ProtoBuff) AddLayer(fieldname string, layer any) {
	// For SubLayer
	if reflect.TypeOf(layer).Kind() == reflect.Struct {
		// Retrieve SubLayer name via GetName
		StructName := reflect.ValueOf(layer).MethodByName("GetName").Call(nil)[0].Interface().(string)

		// Init a LoadedSubLayer struct
		SubLayer := &LoadedSubLayer{LayerName: fieldname, Name: StructName}

		// Build the SubLayer
		StructBuffer := reflect.ValueOf(layer).MethodByName("Build").Call(nil)[0].Interface().(*ProtoBuff)

		// Add SubLayer buffer to main Layer buffer
		binary.Write(buffer, binary.BigEndian, StructBuffer.Bytes())

		for _, field := range StructBuffer.loaded_fields {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: field.Name, Value: field.Value, Enum: field.Enum, ParentLayer: *SubLayer})
		}
	}
}

type LoadedSubLayer struct {
	LayerName string
	Name      string
}

type LoadedField struct {
	Name        string
	Value       any
	Enum        any
	ParentLayer LoadedSubLayer
}

func (f *LoadedField) GetName() string {
	return f.Name
}

func (f *LoadedField) GetValue() any {
	return f.Value
}

func (buffer *ProtoBuff) NextUint8() uint8 {
	return uint8(buffer.Next(1)[0])
}

func (buffer *ProtoBuff) NextUint16() uint16 {
	return binary.BigEndian.Uint16(buffer.Next(2))
}

func (buffer *ProtoBuff) NextUint32() uint32 {
	return binary.BigEndian.Uint32(buffer.Next(4))
}

func (buffer *ProtoBuff) NextUint64() uint64 {
	return binary.BigEndian.Uint64(buffer.Next(8))
}

func (buffer *ProtoBuff) NextUint16Le() uint16 {
	return binary.LittleEndian.Uint16(buffer.Next(2))
}

func (buffer *ProtoBuff) NextUint32Le() uint32 {
	return binary.LittleEndian.Uint32(buffer.Next(4))
}

func (buffer *ProtoBuff) NextUint64Le() uint64 {
	return binary.LittleEndian.Uint64(buffer.Next(8))
}

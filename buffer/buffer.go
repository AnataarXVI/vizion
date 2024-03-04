package buffer

import (
	"bytes"
	"encoding/binary"
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

	// For SubLayer
	if reflect.TypeOf(value).Kind() == reflect.Struct {

		// Retrieve SubLayer name via GetName
		StructName := reflect.ValueOf(value).MethodByName("GetName").Call(nil)[0].Interface().(string)

		// Init a LoadedSubLayer struct
		SubLayer := &LoadedSubLayer{LayerName: fieldname, Name: StructName}

		// Build the SubLayer
		StructBuffer := reflect.ValueOf(value).MethodByName("Build").Call(nil)[0].Interface().(*ProtoBuff)

		// Add SubLayer buffer to main Layer buffer
		binary.Write(buffer, binary.BigEndian, StructBuffer.Bytes())

		for _, field := range StructBuffer.loaded_fields {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: field.Name, Value: field.Value, Enum: field.Enum, ParentLayer: *SubLayer})
		}

	} else if reflect.TypeOf(value).Kind() == reflect.String { // If value is string
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
	// For SubLayer
	if reflect.TypeOf(value).Kind() == reflect.Struct {
		// Retrieve SubLayer name via GetName
		StructName := reflect.ValueOf(value).MethodByName("GetName").Call(nil)[0].Interface().(string)

		// Init a LoadedSubLayer struct
		SubLayer := &LoadedSubLayer{LayerName: fieldname, Name: StructName}

		// Build the SubLayer
		StructBuffer := reflect.ValueOf(value).MethodByName("Build").Call(nil)[0].Interface().(*ProtoBuff)

		for _, field := range StructBuffer.loaded_fields {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: field.Name, Value: field.Value, Enum: field.Enum, ParentLayer: *SubLayer})
		}
	} else {

		// If enum and key exist
		if enum != nil && enum != "" {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value, Enum: enum})
		} else {
			buffer.loaded_fields = append(buffer.loaded_fields, LoadedField{Name: fieldname, Value: value})
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

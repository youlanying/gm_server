package packet

import (
	"reflect"
)

type PKT_rawdata []byte

func Decode_rawdata(reader *Packet) (rawdata PKT_rawdata, err error) {
	return reader.ReadRawData()
}

var (
	rawType reflect.Type
)

func init() {
	rawType = reflect.ValueOf(PKT_rawdata{}).Type()
}

// Write-out struct fields with packet writer.
func Pack(tos int16, tbl interface{}, writer *Packet) []byte {
	if writer == nil {
		writer = Writer()
	}

	v := reflect.ValueOf(tbl)

	switch v.Kind() {
	case reflect.Ptr, reflect.Interface:
		v = v.Elem()
	}
	count := v.NumField()

	// Write code.
	if tos != -1 {
		writer.WriteU16(uint16(tos))
	}

	for i := 0; i < count; i++ {
		f := v.Field(i)

		if f.Type() == rawType {
			writer.WriteRawData(f.Interface().(PKT_rawdata))
			continue
		}

		switch f.Type().Kind() {
		case reflect.Slice, reflect.Array:
			writer.WriteU16(uint16(f.Len()))
			for a := 0; a < f.Len(); a++ {
				if _is_primitive(f.Index(a)) {
					_write_primitive(f.Index(a), writer)
				} else {
					elem := f.Index(a).Interface()
					Pack(-1, elem, writer)
				}
			}
		case reflect.Struct:
			Pack(-1, f.Interface(), writer)
		default:
			_write_primitive(f, writer)
		}
	}

	return writer.Data()
}

// Test whether the field is primitive type.
func _is_primitive(f reflect.Value) bool {
	switch f.Type().Kind() {
	case reflect.Bool,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Float32,
		reflect.Float64,
		reflect.String:
		return true
	}
	return false
}

// Write a primitive field.
func _write_primitive(f reflect.Value, writer *Packet) {
	switch f.Type().Kind() {
	case reflect.Bool:
		writer.WriteBool(f.Bool())
	case reflect.Uint8:
		writer.WriteByte(byte(f.Uint()))
	case reflect.Uint16:
		writer.WriteU16(uint16(f.Uint()))
	case reflect.Uint32:
		writer.WriteU32(uint32(f.Uint()))
	case reflect.Uint64:
		writer.WriteU64(uint64(f.Uint()))
	case reflect.Int:
		writer.WriteU32(uint32(f.Int()))
	case reflect.Int8:
		writer.WriteByte(byte(f.Int()))
	case reflect.Int16:
		writer.WriteU16(uint16(f.Int()))
	case reflect.Int32:
		writer.WriteU32(uint32(f.Int()))
	case reflect.Int64:
		writer.WriteU64(uint64(f.Int()))
	case reflect.Float32:
		writer.WriteFloat32(float32(f.Float()))

	case reflect.Float64:
		writer.WriteFloat64(float64(f.Float()))

	case reflect.String:
		writer.WriteString(f.String())
	}
}

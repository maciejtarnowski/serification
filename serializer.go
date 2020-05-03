package serification

import (
	"errors"
	"fmt"
	"reflect"
)

type TypeSerializer func(Specification) interface{}
type TypeDeserializer func(map[string]interface{}) Specification

type Serializer struct {
	typeSerializers   map[reflect.Type]TypeSerializer
	typeDeserializers map[string]TypeDeserializer
}

func (se *Serializer) Serialize(s Specification) interface{} {
	specType := reflect.TypeOf(s)

	ser, ok := se.typeSerializers[specType]

	if !ok {
		panic(errors.New("type serializer not found for: " + specType.String()))
	}

	return ser(s)
}

func (se *Serializer) Deserialize(d map[string]interface{}) Specification {
	specType, ok := d["type"].(string)
	if !ok {
		panic(errors.New("invalid structure, `type` key not found"))
	}

	deser, ok := se.typeDeserializers[specType]

	if !ok {
		panic(errors.New("type deserializer not found for: " + specType))
	}

	return deser(d)
}

func (se *Serializer) RegisterType(tst reflect.Type, tsFn TypeSerializer, tdt string, tdFn TypeDeserializer) {
	se.typeSerializers[tst] = tsFn
	se.typeDeserializers[tdt] = tdFn
}

func (se *Serializer) RegisterTypeSerializer(typ reflect.Type, fn TypeSerializer) {
	se.typeSerializers[typ] = fn
}

func NewMapSerializer() *Serializer {
	se := Serializer{}

	ts := map[reflect.Type]TypeSerializer{
		reflect.TypeOf(AndSpecification{}): func(s Specification) interface{} {
			spec := s.(AndSpecification)
			data := make(map[string]interface{})

			data["type"] = "and"
			data["left"] = se.Serialize(spec.Left)
			data["right"] = se.Serialize(spec.Right)

			return data
		},
		reflect.TypeOf(OrSpecification{}): func(s Specification) interface{} {
			spec := s.(OrSpecification)
			data := make(map[string]interface{})

			data["type"] = "or"
			data["left"] = se.Serialize(spec.Left)
			data["right"] = se.Serialize(spec.Right)

			return data
		},
		reflect.TypeOf(NotSpecification{}): func(s Specification) interface{} {
			spec := s.(NotSpecification)
			data := make(map[string]interface{})

			data["type"] = "not"
			data["subject"] = se.Serialize(spec.Subject)

			return data
		},
	}

	ds := map[string]TypeDeserializer{
		"and": func(m map[string]interface{}) Specification {
			return AndSpecification{
				Left:  se.Deserialize(m["left"].(map[string]interface{})),
				Right: se.Deserialize(m["right"].(map[string]interface{})),
			}
		},
		"or": func(m map[string]interface{}) Specification {
			return OrSpecification{
				Left:  se.Deserialize(m["left"].(map[string]interface{})),
				Right: se.Deserialize(m["right"].(map[string]interface{})),
			}
		},
		"not": func(m map[string]interface{}) Specification {
			return NotSpecification{
				Subject: se.Deserialize(m["subject"].(map[string]interface{})),
			}
		},
	}

	se.typeSerializers = ts
	se.typeDeserializers = ds

	return &se
}

func NewSQLSerializer() *Serializer {
	se := Serializer{}

	ts := map[reflect.Type]TypeSerializer{
		reflect.TypeOf(AndSpecification{}): func(s Specification) interface{} {
			spec := s.(AndSpecification)

			return fmt.Sprintf("(%s AND %s)", se.Serialize(spec.Left), se.Serialize(spec.Right))
		},
		reflect.TypeOf(OrSpecification{}): func(s Specification) interface{} {
			spec := s.(OrSpecification)

			return fmt.Sprintf("(%s OR %s)", se.Serialize(spec.Left), se.Serialize(spec.Right))
		},
		reflect.TypeOf(NotSpecification{}): func(s Specification) interface{} {
			spec := s.(NotSpecification)

			return fmt.Sprintf("(NOT (%s))", se.Serialize(spec.Subject))
		},
	}

	se.typeSerializers = ts

	return &se
}

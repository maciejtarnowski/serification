package serification

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type TrueSpecification struct{}

func (t TrueSpecification) IsSatisfied(value interface{}) bool {
	return true
}

func (t TrueSpecification) And(other Specification) Specification {
	return AndSpecification{Left: t, Right: other}
}

func (t TrueSpecification) Or(other Specification) Specification {
	return OrSpecification{Left: t, Right: other}
}

func (t TrueSpecification) Not() Specification {
	return NotSpecification{Subject: t}
}

func TestSerializeAndSpecification(t *testing.T) {
	s := NewSerializer()
	s.RegisterType(reflect.TypeOf(TrueSpecification{}), func(s Specification) map[string]interface{} {
		return map[string]interface{}{"type": "true"}
	}, "true", func(m map[string]interface{}) Specification {
		return TrueSpecification{}
	})

	spec := TrueSpecification{}.And(TrueSpecification{})

	expected := map[string]interface{}{
		"type": "and",
		"left": map[string]interface{}{
			"type": "true",
		},
		"right": map[string]interface{}{
			"type": "true",
		},
	}

	assert.Equal(t, expected, s.Serialize(spec))
}

func TestSerializeOrSpecification(t *testing.T) {
	s := NewSerializer()
	s.RegisterType(reflect.TypeOf(TrueSpecification{}), func(s Specification) map[string]interface{} {
		return map[string]interface{}{"type": "true"}
	}, "true", func(m map[string]interface{}) Specification {
		return TrueSpecification{}
	})

	spec := TrueSpecification{}.Or(TrueSpecification{})

	expected := map[string]interface{}{
		"type": "or",
		"left": map[string]interface{}{
			"type": "true",
		},
		"right": map[string]interface{}{
			"type": "true",
		},
	}

	assert.Equal(t, expected, s.Serialize(spec))
}

func TestSerializeNotSpecification(t *testing.T) {
	s := NewSerializer()
	s.RegisterType(reflect.TypeOf(TrueSpecification{}), func(s Specification) map[string]interface{} {
		return map[string]interface{}{"type": "true"}
	}, "true", func(m map[string]interface{}) Specification {
		return TrueSpecification{}
	})

	spec := TrueSpecification{}.Not()

	expected := map[string]interface{}{
		"type": "not",
		"subject": map[string]interface{}{
			"type": "true",
		},
	}

	assert.Equal(t, expected, s.Serialize(spec))
}

func TestSerializeComplexSpecification(t *testing.T) {
	s := NewSerializer()
	s.RegisterType(reflect.TypeOf(TrueSpecification{}), func(s Specification) map[string]interface{} {
		return map[string]interface{}{"type": "true"}
	}, "true", func(m map[string]interface{}) Specification {
		return TrueSpecification{}
	})

	spec := TrueSpecification{}.Or(TrueSpecification{}.And(TrueSpecification{}.Not()))

	expected := map[string]interface{}{
		"type": "or",
		"left": map[string]interface{}{
			"type": "true",
		},
		"right": map[string]interface{}{
			"type": "and",
			"left": map[string]interface{}{
				"type": "true",
			},
			"right": map[string]interface{}{
				"type": "not",
				"subject": map[string]interface{}{
					"type": "true",
				},
			},
		},
	}

	assert.Equal(t, expected, s.Serialize(spec))
}

func TestDeserializeAndSpecification(t *testing.T) {
	s := NewSerializer()
	s.RegisterType(reflect.TypeOf(TrueSpecification{}), func(s Specification) map[string]interface{} {
		return map[string]interface{}{"type": "true"}
	}, "true", func(m map[string]interface{}) Specification {
		return TrueSpecification{}
	})

	expected := TrueSpecification{}.And(TrueSpecification{})

	data := map[string]interface{}{
		"type": "and",
		"left": map[string]interface{}{
			"type": "true",
		},
		"right": map[string]interface{}{
			"type": "true",
		},
	}

	assert.Equal(t, expected, s.Deserialize(data))
}

func TestDeserializeOrSpecification(t *testing.T) {
	s := NewSerializer()
	s.RegisterType(reflect.TypeOf(TrueSpecification{}), func(s Specification) map[string]interface{} {
		return map[string]interface{}{"type": "true"}
	}, "true", func(m map[string]interface{}) Specification {
		return TrueSpecification{}
	})

	expected := TrueSpecification{}.Or(TrueSpecification{})

	data := map[string]interface{}{
		"type": "or",
		"left": map[string]interface{}{
			"type": "true",
		},
		"right": map[string]interface{}{
			"type": "true",
		},
	}

	assert.Equal(t, expected, s.Deserialize(data))
}

func TestDeserializeNotSpecification(t *testing.T) {
	s := NewSerializer()
	s.RegisterType(reflect.TypeOf(TrueSpecification{}), func(s Specification) map[string]interface{} {
		return map[string]interface{}{"type": "true"}
	}, "true", func(m map[string]interface{}) Specification {
		return TrueSpecification{}
	})

	expected := TrueSpecification{}.Not()

	data := map[string]interface{}{
		"type": "not",
		"subject": map[string]interface{}{
			"type": "true",
		},
	}

	assert.Equal(t, expected, s.Deserialize(data))
}

func TestDeserializeComplexSpecification(t *testing.T) {
	s := NewSerializer()
	s.RegisterType(reflect.TypeOf(TrueSpecification{}), func(s Specification) map[string]interface{} {
		return map[string]interface{}{"type": "true"}
	}, "true", func(m map[string]interface{}) Specification {
		return TrueSpecification{}
	})

	expected := TrueSpecification{}.Or(TrueSpecification{}.And(TrueSpecification{}.Not()))

	data := map[string]interface{}{
		"type": "or",
		"left": map[string]interface{}{
			"type": "true",
		},
		"right": map[string]interface{}{
			"type": "and",
			"left": map[string]interface{}{
				"type": "true",
			},
			"right": map[string]interface{}{
				"type": "not",
				"subject": map[string]interface{}{
					"type": "true",
				},
			},
		},
	}

	assert.Equal(t, expected, s.Deserialize(data))
}

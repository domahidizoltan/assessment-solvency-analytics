package businesstask

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type Type string

const (
	TypeString Type = "string"
	TypeInt    Type = "integer"
	TypeBool   Type = "boolean"
)

var (
	ErrMissingRequiredKey       = errors.New("required key is missing")
	ErrUnexpectedKey            = errors.New("unexpected key")
	ErrUnmarshalEnvelope        = errors.New("failed to unmarshal envelope")
	ErrUnexpectedType           = errors.New("unexpected type")
	ErrInvalidType              = errors.New("invalid type")
	ErrUnexpectedEnvelopeKey    = errors.New("unexpected envelope key")
	ErrMissingType              = errors.New("missing type")
	ErrUnexpectedSchemaProperty = errors.New("unexpected schema property")

	validTypes = map[Type]struct{}{
		TypeString: {},
		TypeInt:    {},
		TypeBool:   {},
	}
)

type (
	SchemaProperties struct {
		Type     string `json:"type,omitempty"`
		Required bool   `json:"required,omitempty"`
	}

	Envelope struct {
		Schema   map[string]SchemaProperties `json:"schema"`
		Document map[string]any              `json:"document"`
	}
)

func Validate(input []byte, forceTypeValidation bool) error {
	var envelope Envelope
	if err := json.Unmarshal(input, &envelope); err != nil {
		return wrapErr(ErrUnmarshalEnvelope, err.Error())
	}

	for key, properties := range envelope.Schema {
		if properties.Required {
			if _, ok := envelope.Document[key]; !ok {
				return wrapErr(ErrMissingRequiredKey, key)
			}
		}

		if !isValidType(properties.Type, forceTypeValidation) {
			return wrapErr(ErrInvalidType, properties.Type)
		}
	}

	for key, val := range envelope.Document {
		var properties SchemaProperties
		if props, ok := envelope.Schema[key]; !ok {
			return wrapErr(ErrUnexpectedKey, key)
		} else {
			properties = props
		}

		if !isExpectedFieldType(properties.Type, val, forceTypeValidation) {
			return wrapErr(ErrUnexpectedType, fmt.Sprintf("key=%s type=%s", key, properties.Type))
		}
	}
	return nil
}

func wrapErr(err error, detail string) error {
	return fmt.Errorf("%w: %s", err, detail)
}

func isValidType(t string, forceValidation bool) bool {
	if !forceValidation || len(t) == 0 {
		return true
	}

	_, ok := validTypes[Type(t)]
	return ok
}

func isExpectedFieldType(expectedType string, val any, forceValidation bool) bool {
	if !forceValidation && len(expectedType) == 0 {
		return true
	}

	switch reflect.TypeOf(val).Name() {
	case "string":
		return expectedType == string(TypeString)
	case "bool":
		return expectedType == string(TypeBool)
	case "float64", "float32", "int", "int64", "int32", "uint", "uint8", "uint16", "uint32", "uint64":
		return expectedType == string(TypeInt)
	default:
		return false
	}

}

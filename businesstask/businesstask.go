package businesstask

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Type string

const (
	TypeString Type = "string"
	TypeInt    Type = "integer"
	TypeBool   Type = "boolean"
)

var (
	ErrMissingRequiredKey = errors.New("required key is missing")
	ErrUnexpectedField    = errors.New("unexpected field")
	ErrUnmarshalEnvelope  = errors.New("failed to unmarshal envelope")
	ErrUnexpectedType     = errors.New("unexpected type")
	ErrInvalidType        = errors.New("invalid type")
	ErrUnexpectedKey      = errors.New("unexpected key")
	ErrMissingType        = errors.New("missing type")

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
	envelope, err := getEnvelope(input)
	if err != nil {
		return err
	}

	for key, properties := range envelope.Schema {
		if properties.Required {
			if _, ok := envelope.Document[key]; !ok {
				return wrapErr(ErrMissingRequiredKey, key)
			}
		}

		if err := validateType(properties.Type, forceTypeValidation); err != nil {
			return err
		}
	}

	for key, val := range envelope.Document {
		var properties SchemaProperties
		if props, ok := envelope.Schema[key]; !ok {
			return wrapErr(ErrUnexpectedField, key)
		} else {
			properties = props
		}

		if !isExpectedFieldType(properties.Type, val, forceTypeValidation) {
			return wrapErr(ErrUnexpectedType, fmt.Sprintf("key=%s type=%s", key, properties.Type))
		}
	}
	return nil
}

func getEnvelope(input []byte) (*Envelope, error) {
	var envelope Envelope

	decoder := json.NewDecoder(bytes.NewReader(input))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&envelope); err != nil {
		if strings.HasPrefix(err.Error(), "json: unknown field") {
			return nil, wrapErr(ErrUnexpectedKey, err.Error())
		}

		return nil, wrapErr(ErrUnmarshalEnvelope, err.Error())
	}
	return &envelope, nil
}

func wrapErr(err error, detail string) error {
	return fmt.Errorf("%w: %s", err, detail)
}

func validateType(t string, forceValidation bool) error {
	if !forceValidation {
		return nil
	} else if len(t) == 0 {
		return wrapErr(ErrMissingType, t)
	}

	if _, ok := validTypes[Type(t)]; !ok {
		return wrapErr(ErrInvalidType, t)
	}

	return nil
}

func isExpectedFieldType(expectedType string, val any, forceValidation bool) bool {
	if !forceValidation && len(expectedType) == 0 {
		return true
	}

	switch reflect.TypeOf(val).String() {
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

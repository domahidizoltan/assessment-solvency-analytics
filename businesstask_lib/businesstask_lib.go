package businesstask_lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
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

	newLoader = gojsonschema.NewGoLoader
)

type (
	SchemaProperties struct {
		Type     string `json:"type,omitempty"`
		Required bool   `json:"required,omitempty"`
	}

	Envelope struct {
		Schema   map[string]map[string]any `json:"schema"`
		Required []string                  `json:"-"`
		Document json.RawMessage           `json:"document"`
	}
)

func Validate(input []byte, forceTypeValidation bool) error {
	envelope, err := getEnvelope(input, forceTypeValidation)
	if err != nil {
		return err
	}

	schemaLoader := newLoader(map[string]any{
		"type":                 "object",
		"properties":           envelope.Schema,
		"required":             envelope.Required,
		"additionalProperties": false,
	})

	schema, err := gojsonschema.NewSchema(schemaLoader)
	if err != nil {
		return err
	}

	documentLoader := newLoader(envelope.Document)

	result, err := schema.Validate(documentLoader)
	if err != nil {
		return err
	}
	if !result.Valid() {
		return errors.New(result.Errors()[0].Description())
	}

	return nil
}

func getEnvelope(input []byte, forceTypeValidation bool) (*Envelope, error) {
	var envelope Envelope

	decoder := json.NewDecoder(bytes.NewReader(input))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&envelope); err != nil {
		if strings.HasPrefix(err.Error(), "json: unknown field") {
			return nil, wrapErr(ErrUnexpectedKey, err.Error())
		}

		return nil, wrapErr(ErrUnmarshalEnvelope, err.Error())
	}

	envelope.Required = []string{}
	for field, def := range envelope.Schema {
		for k := range def {
			if k != "required" && k != "type" {
				return nil, wrapErr(ErrUnexpectedKey, k)
			}
		}
		if required, ok := def["required"].(bool); ok {
			if required {
				envelope.Required = append(envelope.Required, field)
			}
			delete(def, "required")
		}
		if t, ok := def["type"]; ok {
			if err := validateType(t.(string), forceTypeValidation); err != nil {
				return nil, err
			}
		} else if forceTypeValidation {
			return nil, wrapErr(ErrMissingType, field)
		}
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

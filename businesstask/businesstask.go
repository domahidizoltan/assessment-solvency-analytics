package businesstask

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrMissingRequiredKey = errors.New("required key is missing")
	ErrUnexpectedKey      = errors.New("unexpected key")
	ErrUnmarshalEnvelope  = errors.New("failed to unmarshal envelope")
)

type (
	SchemaProperties struct {
		Required bool `json:"required"`
	}

	Envelope struct {
		Schema   map[string]SchemaProperties `json:"schema"`
		Document map[string]any              `json:"document"`
	}
)

func Validate(input []byte) error {
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
	}

	for key := range envelope.Document {
		if _, ok := envelope.Schema[key]; !ok {
			return wrapErr(ErrUnexpectedKey, key)
		}
	}
	return nil
}

func wrapErr(err error, detail string) error {
	return fmt.Errorf("%w: %s", err, detail)
}

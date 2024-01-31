package businesstask

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	dir      = "../testdata/"
	fileType = ".json"

	fixture_0_0_valid1                              = dir + "0_0_valid1" + fileType
	fixture_0_1_valid2_optional_field_is_missing    = dir + "0_1_valid2_optional_field_is_missing" + fileType
	fixture_0_2_invalid_document_key1_is_missing    = dir + "0_2_invalid_document-key1_is_missing" + fileType
	fixture_0_3_invalid_document_key3_is_unexpected = dir + "0_3_invalid_document-key3_is_unexpected" + fileType

	fixture_1_0_valid                                 = dir + "1_0_valid" + fileType
	fixture_1_1_invalid_document_key1_is_not_a_string = dir + "1_1_invalid_document-key1_is_not_a_string" + fileType

	fixture_2_0_invalid_unexpected_key_something_else                             = dir + "2_0_invalid_unexpected_key_something_else" + fileType
	fixture_2_1_invalid_incomplete_schema_key1_type_is_missing                    = dir + "2_1_invalid_incomplete_schema-key1-type_is_missing" + fileType
	fixture_2_2_invalid_invalid_schema_unexpected_key_schema__key1_something_else = dir + "2_2_invalid_invalid_schema_unexpected_key_schema-key1-something_else" + fileType
)

func TestValidate(t *testing.T) {
	for _, s := range []struct {
		name, fixtureFile, fixtureString string
		expectedError                    error
	}{
		{
			name:        "valid1",
			fixtureFile: fixture_0_0_valid1,
		},
		{
			name:        "valid2_optional_field_is_missing",
			fixtureFile: fixture_0_1_valid2_optional_field_is_missing,
		},
		{
			name:          "invalid_document_key1_is_missing",
			fixtureFile:   fixture_0_2_invalid_document_key1_is_missing,
			expectedError: ErrMissingRequiredKey,
		},
		{
			name:          "invalid_document_key3_is_unexpected",
			fixtureFile:   fixture_0_3_invalid_document_key3_is_unexpected,
			expectedError: ErrUnexpectedKey,
		},
		{
			name:          "envelope_unmarshal_error",
			fixtureString: "{",
			expectedError: ErrUnmarshalEnvelope,
		},
		{
			name:        "valid_types",
			fixtureFile: fixture_1_0_valid,
		},
		{
			name:          "invalid_document_key1_is_not_a_string",
			fixtureFile:   fixture_1_1_invalid_document_key1_is_not_a_string,
			expectedError: ErrUnexpectedType,
		},
		{
			name: "invalid_type",
			fixtureString: `
			{
				"schema": {
					"key1": {
						"type": "array"
					}
				},
				"document": {}
			}`,
			expectedError: ErrInvalidType,
		},
		{
			name:          "invalid_unexpected_key_something_else",
			fixtureFile:   fixture_2_0_invalid_unexpected_key_something_else,
			expectedError: ErrUnexpectedEnvelopeKey,
		},
		{
			name:          "invalid_incomplete_schema_key1_type_is_missing",
			fixtureFile:   fixture_2_1_invalid_incomplete_schema_key1_type_is_missing,
			expectedError: ErrMissingType,
		},
		{
			name:          "invalid_invalid_schema_unexpected_key_schema_key1_something_else",
			fixtureFile:   fixture_2_2_invalid_invalid_schema_unexpected_key_schema__key1_something_else,
			expectedError: ErrUnexpectedSchemaProperty,
		},
	} {
		var fixture []byte
		if len(s.fixtureString) > 0 {
			fixture = []byte(s.fixtureString)
		} else {
			var err error
			fixture, err = os.ReadFile(s.fixtureFile)
			require.NoError(t, err)
		}

		actualErr := Validate(fixture)
		assert.ErrorIs(t, actualErr, s.expectedError)
	}

}

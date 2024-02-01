package businesstask_lib

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonreference"
	"github.com/xeipuuv/gojsonschema"
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
		forceTypeValidation              bool
		expectedError                    string
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
			expectedError: "key1 is required",
		},
		{
			name:          "invalid_document_key3_is_unexpected",
			fixtureFile:   fixture_0_3_invalid_document_key3_is_unexpected,
			expectedError: "Additional property key3 is not allowed",
		},
		{
			name:          "envelope_unmarshal_error",
			fixtureString: "{",
			expectedError: ErrUnmarshalEnvelope.Error(),
		},
		{
			name:                "valid_types",
			fixtureFile:         fixture_1_0_valid,
			forceTypeValidation: true,
		},
		{
			name:                "invalid_document_key1_is_not_a_string",
			fixtureFile:         fixture_1_1_invalid_document_key1_is_not_a_string,
			forceTypeValidation: true,
			expectedError:       "Invalid type. Expected: string, given: integer",
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
				"document": {
					"key1": []
				}
			}`,
			forceTypeValidation: true,
			expectedError:       ErrInvalidType.Error(),
		},
		{
			name:                "invalid_unexpected_key_something_else",
			fixtureFile:         fixture_2_0_invalid_unexpected_key_something_else,
			forceTypeValidation: true,
			expectedError:       ErrUnexpectedKey.Error(),
		},
		{
			name:                "invalid_incomplete_schema_key1_type_is_missing",
			fixtureFile:         fixture_2_1_invalid_incomplete_schema_key1_type_is_missing,
			forceTypeValidation: true,
			expectedError:       ErrMissingType.Error(),
		},
		{
			name:                "invalid_invalid_schema_unexpected_key_schema_key1_something_else",
			fixtureFile:         fixture_2_2_invalid_invalid_schema_unexpected_key_schema__key1_something_else,
			forceTypeValidation: true,
			expectedError:       ErrUnexpectedKey.Error(),
		},
		{
			name: "schema_loader_error",
			fixtureString: `
			{
				"schema": {
					"key1": {
						"type": "str"
					}
				}
			}`,
			expectedError: "has a primitive type that is NOT VALID -- given: /str/ Expected valid values are:[array boolean integer number null object string]",
		},
		{
			name: "document_loader_error",
			fixtureString: `
			{
				"schema": {},
				"document": "FAIL"
			}`,
			expectedError: "loader failed",
		},
		{
			name: "type_validation_forced",
			fixtureString: `
			{
				"schema": {
					"key1": {
						"type": ""
					}
				}
			}`,
			forceTypeValidation: true,
			expectedError:       ErrMissingType.Error(),
		},
	} {
		t.Run(s.name, func(t *testing.T) {
			if s.name == "document_loader_error" {
				bkp := newLoader
				defer func() {
					newLoader = bkp
				}()
				newLoader = func(source any) gojsonschema.JSONLoader {
					if reflect.TypeOf(source).String() == "json.RawMessage" {
						return fakeJSONLoader{}
					}
					return gojsonschema.NewGoLoader(source)
				}
			}

			var fixture []byte
			if len(s.fixtureString) > 0 {
				fixture = []byte(s.fixtureString)
			} else {
				var err error
				fixture, err = os.ReadFile(s.fixtureFile)
				require.NoError(t, err)
			}

			actualErr := Validate(fixture, s.forceTypeValidation)
			if len(s.expectedError) == 0 {
				assert.NoError(t, actualErr)
			} else {
				assert.ErrorContains(t, actualErr, s.expectedError)
			}
		})
	}
}

func BenchmarkValidate(b *testing.B) {
	fixture, err := os.ReadFile(fixture_1_0_valid)
	require.NoError(b, err)

	for i := 0; i < b.N; i++ {
		Validate(fixture, true)
	}
}

type fakeJSONLoader struct{}

func (fakeJSONLoader) JsonSource() interface{} { return nil }
func (fakeJSONLoader) LoadJSON() (interface{}, error) {
	return nil, errors.New("loader failed")
}
func (fakeJSONLoader) JsonReference() (gojsonreference.JsonReference, error) {
	return gojsonreference.JsonReference{}, nil
}
func (fakeJSONLoader) LoaderFactory() gojsonschema.JSONLoaderFactory { return nil }

package businesstask

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	dir                                             = "../testdata/"
	fixture_0_0_valid1                              = dir + "0_0_valid1.json"
	fixture_0_1_valid2_optional_field_is_missing    = dir + "0_1_valid2_optional_field_is_missing.json"
	fixture_0_2_invalid_document_key1_is_missing    = dir + "0_2_invalid_document-key1_is_missing.json"
	fixture_0_3_invalid_document_key3_is_unexpected = dir + "0_3_invalid_document-key3_is_unexpected.json"
)

func TestValidateSimple(t *testing.T) {
	for _, s := range []struct {
		name          string
		expectedError error
	}{
		{
			name: "valid1",
		},
		{
			name: "valid2_optional_field_is_missing",
		},
		{
			name:          "invalid_document_key1_is_missing",
			expectedError: ErrMissingRequiredKey,
		},
		{
			name:          "invalid_document_key3_is_unexpected",
			expectedError: ErrUnexpectedKey,
		},
	} {
		fixture, err := os.ReadFile(fixture_0_0_valid1)
		require.NoError(t, err)

		actualErr := Validate(string(fixture))
		assert.ErrorIs(t, actualErr, s.expectedError)
	}

}

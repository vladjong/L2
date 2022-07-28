package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpackingString(t *testing.T) {
	input := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
	}

	expected := []string{
		"aaaabccddddde",
		"abcd",
		"",
		"",
	}

	for i, data := range input {
		result, err := unpackingString(data)
		assert.NoError(t, err)
		assert.Equal(t, expected[i], result)
	}
}

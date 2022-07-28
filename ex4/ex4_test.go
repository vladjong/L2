package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestsearchAnagramFirst(t *testing.T) {
	input := []string{
		"слиток",
		"автобус",
		"пятка",
		"Столик",
		"Столик",
		"тяпка",
	}

	expected := map[string][]string{
		"пятка" : {"пятка", "тяпка"},
		"слиток:" : {"слиток", "столик"}
	}

	for i, data := range input {
		result := searchAnagram(data)
		assert.Equal(t, expected, result)
	}
}

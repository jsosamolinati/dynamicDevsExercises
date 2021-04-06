package DDATM_ES_0001_test

import (
	"testing"

	"github.com/jsosamolinati/dynamicDevsExercises/DDATM_ES_0001"
	"github.com/stretchr/testify/assert"
)

func Test_retrieveOrderStringWithBlankSpace(t *testing.T) {
	result := DDATM_ES_0001.Exec("123 a")

	assert.Equal(t, result, "123 b")
}

func Test_retrieveOrderStringWithUpperLetter(t *testing.T) {
	result := DDATM_ES_0001.Exec("123Abcf4")

	assert.Equal(t, result, "123Bcdg4")
}

func Test_retrieveOrderStringWithZ(t *testing.T) {
	result := DDATM_ES_0001.Exec("123 z")

	assert.Equal(t, result, "123 a")
}
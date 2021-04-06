package exercises_test

import (
	"testing"

	"github.com/jsosamolinati/dynamicDevsExercises/exercises"
	"github.com/stretchr/testify/assert"
)

func Test_retrieveOrderStringWithBlankSpace(t *testing.T) {
	result := exercises.Exec("123 a")

	assert.Equal(t, result, "123 b")
}

func Test_retrieveOrderStringWithUpperLetter(t *testing.T) {
	result := exercises.Exec("123Abcf4")

	assert.Equal(t, result, "123Bcdg4")
}

func Test_retrieveOrderStringWithZ(t *testing.T) {
	result := exercises.Exec("123 z")

	assert.Equal(t, result, "123 a")
}
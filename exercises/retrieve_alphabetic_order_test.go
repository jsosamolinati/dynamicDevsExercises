package exercises_test

import (
	"testing"

	"github.com/jsosamolinati/dynamicDevsExercises/exercises"
	"github.com/stretchr/testify/assert"
)

func test_retrieveOrderString(t *testing.T) {
	result := exercises.Exec("123a")

	assert.Equal(t, result, "123b")
}

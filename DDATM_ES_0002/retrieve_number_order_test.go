package DDATM_ES_0002_test

import (
	"testing"

	"github.com/jsosamolinati/dynamicDevsExercises/DDATM_ES_0002"
	"github.com/stretchr/testify/assert"
)

func Test_retrieveOrderNumbers(t *testing.T) {
	mockParam := []int{5, 7, 3, 9}
	excpected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result, err := DDATM_ES_0002.Exec(mockParam)

	assert.Equal(t, result, excpected)
	assert.Nil(t, err)
}

func Test_retrieveOrderNumbersAlternative(t *testing.T) {
	mockParam := []int{5, 7, 3, 9}
	excpected := []int{3, 4, 5, 6, 7, 8, 9}
	result, err := DDATM_ES_0002.ExecAlternative(mockParam)

	assert.Equal(t, result, excpected)
	assert.Nil(t, err)
}
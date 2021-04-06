package exercises

import (
	"testing"
	"retrieve_alphabetic_order"
	"github.com/stretchr/testify/assert"
)

func test_retrieveOrderString(t *testing.T) {
	result := retrieve_alphabetic_order.Execute("123a")

   assert.Equal(result, "123b")
}
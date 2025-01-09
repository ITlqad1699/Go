package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	// result := AddOne(1)
	// if result != 2 {
	// 	t.Errorf("Expected 2 but got %d", result)
	// }
	assert.Equal(t, AddOne(1), 2, "Expected 2 but got 2")
}

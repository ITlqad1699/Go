package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// func TestAddOne(t *testing.T) {
// 	// result := AddOne(1)
// 	// if result != 2 {
// 	// 	t.Errorf("Expected 2 but got %d", result)
// 	// }
// 	assert.Equal(t, AddOne(1), 2, "Expected 2 but got 2")
// }

// Test Assert
func TestAssert(t *testing.T) {
	assert.Equal(t, AddOne(1), 2, "Expected 2 but got 2")
	fmt.Println("executing")
}

// Test Require
func TestRequire(t *testing.T) {
	require.Equal(t, AddOne(1), 2, "Expected 2 but got 2")
	fmt.Println("not executing")
}

func AddOne(x int) int {
	return x + 1
}

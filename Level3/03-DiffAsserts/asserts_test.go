/*
>>> Different Asserts

Statement
Using Testify/assert pkg, execute different functions and validate the expected result. 
You can create as many cases as you want
Equal | NotEqual | Nil | NotNil

Topics to Practice:
testing, multiple cases, testify
*/
package main

import(
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestWithTestify(t *testing.T) {
	assert.Equal(t, 1, 1, "They should be equal")
	assert.Equal(t, 2, 2, "They should be equal")

	assert.NotEqual(t, 3, 2, "They should be different")
	assert.NotEqual(t, 2, 1, "They should be different")

	assert.Nil(t, nil)

	var object = 123

	if assert.NotNil(t, object) {
		assert.NotEqual(t, 6, 6, "They should be different")
	}

	assert.Nil(t, object)

	var arr1 = []int{2, 3, 5, 4}
	var arr2 = []int{2, 3, 5}
	assert.Equal(t, arr1, arr2, "err")
}
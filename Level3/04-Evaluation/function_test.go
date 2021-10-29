package main
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFilterByAgeRange(t *testing.T){
	num1:=6
	num2:=13
	var slc = []int{2, 3, 5, 7, 11, 13}

	v := filterByAgeRange(num1, num2, slc)

	var want =  []int{7, 11, 1}
	assert.Equal(t, v, want, "Array is not equals than expected")
}
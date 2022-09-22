package example1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	validData = Data{
		Items: []Item{
			{1, "one"},
			{2, "two"},
			{3, "three"},
			{4, "four"},
			{5, "five"},
		}}
)

func TestDec(t *testing.T) {
	t.Run("plain decode", func(tt *testing.T) {
		res, err := plainDecode()
		assert.Nil(tt, err)
		assert.EqualValues(tt, &validData, res)
	})

	t.Run("decode with any, no constraint", func(tt *testing.T) {
		res, err := decodeWithAny[Data]()
		assert.Nil(tt, err)
		assert.EqualValues(tt, &validData, res)
	})

	t.Run("decode with constraint", func(tt *testing.T) {
		res, err := decodeWithConstraint[*Data]()
		assert.Nil(tt, err)
		assert.EqualValues(tt, &validData, *res)
	})

	t.Run("decode with constraint and method", func(tt *testing.T) {
		res, err := decodeWithConstraintAndMethod[Data]()
		assert.Nil(tt, err)
		assert.EqualValues(tt, &validData, res)
	})
}

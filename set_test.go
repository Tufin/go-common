package common

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestArray_AddItems(t *testing.T) {

	arr := NewSet().AddItems([]string{"a", "c"}).Add("b")
	assert.True(t, arr.IsSimilar([]string{"a", "b", "c"}))
	assert.Equal(t, 3, arr.Size())
}
func TestArray_Contains(t *testing.T) {

	arr := NewSet().AddItems([]string{"b", "a"})
	assert.True(t, arr.Contains("a"))
	assert.True(t, arr.Contains("b"))
	assert.False(t, arr.Contains("c"))
}

func TestArray_IsSimilar_True(t *testing.T) {

	arr := NewSet()
	arr.AddItems([]string{"a", "b", "c"})
	assert.True(t, arr.IsSimilar([]string{"b", "c", "a"}))
}

func TestArray_Size(t *testing.T) {

	arr := NewSet()
	arr.AddItems([]string{"a", "b", "c"})
	assert.Equal(t, 3, arr.Size())
}

func TestArray_Items(t *testing.T) {

	const name = "nehmad"
	list := NewSet().Add(name)
	assert.Equal(t, 1, len(list.Items()))
	assert.True(t, list.Contains(name))
}

func TestArray_ItemsEmpty(t *testing.T) {

	assert.Equal(t, 0, len(NewSet().Items()))
}

func TestArray_IsSimilarEmptyArrays(t *testing.T) {

	arr := NewSet()
	assert.True(t, arr.IsSimilar([]string{}))
}

func TestArray_IsSimilarSize_False(t *testing.T) {

	arr := NewSet()
	arr.AddItems([]string{"a", "c"})
	assert.False(t, arr.IsSimilar([]string{"b", "c", "a"}))
}

func TestArray_IsSimilar_False(t *testing.T) {

	arr := NewSet()
	arr.AddItems([]string{"a", "c", "d"})
	assert.False(t, arr.IsSimilar([]string{"b", "c", "a"}))
}

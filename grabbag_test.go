package grabbag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrabBagFromData(t *testing.T) {
	t.Parallel()
	gb := FromData(sampleData())
	assert.NotNil(t, gb, "FromData gave us nil")
}

func TestGrabBag_Grab(t *testing.T) {
	t.Parallel()
	gb := FromData(sampleData())
	result := gb.Grab("hello")
	assert.NotNil(t, result, "Grab returned nil")
}

func TestGrabBag_Has(t *testing.T) {
	t.Parallel()
	gb := FromData(sampleData())
	result := gb.Has("hello")
	assert.Equal(t, result, true, "Grab bag returned a false negative for has")

	result = gb.Has("hello.someOther.field")
	assert.Equal(t, result, false, "Grab bag returned a false positive for has")

	result = gb.Has("nested.inner.another.value")
	assert.Equal(t, result, true, "Grab bag returned a false negative for has on a nested field")
}

func TestGrabBag_Nested(t *testing.T) {
	t.Parallel()
	gb := FromData(sampleData())
	result := gb.String("nested.hello")
	assert.Equal(t, "world inner", result, "Grab bag did not return proper string")

	i := gb.Int("nested.inner.another.value")
	assert.Equal(t, i, 5, "Grab bag did not return a proper value for a nested int")
}

func TestGrabBag_Types(t *testing.T) {
	t.Parallel()
	gb := FromData(sampleData())

	str := gb.String("hello")
	assert.Equal(t, "world", str, "Grab bag did not return proper string")

	i := gb.Int("int")
	assert.Equal(t, 1, i, "Grab bag did not return proper int")

	f32 := gb.Float32("float32")
	assert.Equal(t, float32(1), f32, "Grab bag did not return proper float32")

	f64 := gb.Float64("float64")
	assert.Equal(t, float64(1), f64, "Grab bag did not return proper float64")

	b := gb.Bool("bool")
	assert.Equal(t, true, b, "Grab bag did not return proper bool")

	intSlice := gb.IntSlice("intSlice")
	assert.Len(t, intSlice, 5, "Grab bag missed some ints in our slice")
	assert.Equal(t, intSlice, []int{1, 2, 3, 4, 5}, "Grab bag did not return proper int slice")

	strSlice := gb.StringSlice("stringSlice")
	assert.Len(t, strSlice, 2, "Grab bag missed some strings in our slice")
	assert.Equal(t, strSlice, []string{"hello", "world"}, "Grab bag did not return proper string slice")
}

func BenchmarkGrabBag_Simple(b *testing.B) {
	gb := FromData(sampleData())

	for i := 0; i < b.N; i++ {
		gb.Grab("hello")
	}
}

func BenchmarkGrabBag_Traverse(b *testing.B) {
	gb := FromData(sampleData())

	for i := 0; i < b.N; i++ {
		gb.Grab("nested.inner.another.value")
	}
}

func sampleData() map[string]interface{} {
	return map[string]interface{}{
		"hello":       "world",
		"int":         1,
		"float32":     float32(1),
		"float64":     float64(1),
		"bool":        true,
		"intSlice":    []int{1, 2, 3, 4, 5},
		"stringSlice": []string{"hello", "world"},
		"nested": map[string]interface{}{
			"hello": "world inner",
			"inner": map[string]interface{}{
				"another": map[string]interface{}{
					"value": 5,
				},
			},
		},
	}
}

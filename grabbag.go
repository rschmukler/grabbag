package grabbag

import (
	"strings"
)

// GrabBag is the only type in GrabBag. It is an API to traverse nested map data
type GrabBag struct {
	data map[string]interface{}
}

// Creates a new GrabBag from raw data
func FromData(data map[string]interface{}) *GrabBag {
	return &GrabBag{data}
}

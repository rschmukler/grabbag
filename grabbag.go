package grabbag

import "strings"

// Seperator is the seperator used by grabbag to traverse
var Seperator = "."

// GrabBag is the only type in GrabBag. It is an API to traverse nested map data
type GrabBag struct {
	data interface{}
}

// FromData creates a new GrabBag from raw data
func FromData(data interface{}) *GrabBag {
	return &GrabBag{data}
}

// Grab gets a generic interface from our data
func (g *GrabBag) Grab(query string) interface{} {
	levels := strings.Split(query, Seperator)
	path := levels[0]
	if asMap, ok := g.data.(map[string]interface{}); ok {
		if len(levels) > 1 {
			return FromData(asMap[path]).Grab(strings.Join(levels[1:], Seperator))
		}
		return asMap[path]
	}
	return nil
}

// Has returns whether our data has the selected query available
func (g *GrabBag) Has(query string) bool {
	levels := strings.Split(query, Seperator)
	path := levels[0]
	if asMap, ok := g.data.(map[string]interface{}); ok {
		if len(levels) > 1 {
			return FromData(asMap[path]).Has(strings.Join(levels[1:], Seperator))
		}
		return asMap[path] != nil
	}
	return false
}

// String grabs a query casted as a string
func (g *GrabBag) String(query string) string {
	res := g.Grab(query)
	if res != nil {
		return res.(string)
	}
	return ""
}

// StringSlice grabs a query casted as a []string
func (g *GrabBag) StringSlice(query string) []string {
	res := g.Grab(query)
	if res != nil {
		return res.([]string)
	}
	return []string{}
}

// Int grabs a query casted as an int
func (g *GrabBag) Int(query string) int {
	res := g.Grab(query)
	if res != nil {
		return res.(int)
	}
	return 0
}

// IntSlice grabs a query casted as an []int
func (g *GrabBag) IntSlice(query string) []int {
	res := g.Grab(query)
	if res != nil {
		return res.([]int)
	}
	return []int{}
}

// Bool grabs a query casted as a bool
func (g *GrabBag) Bool(query string) bool {
	res := g.Grab(query)
	if res != nil {
		return res.(bool)
	}
	return false
}

// Float32 grabs a query casted as a float32
func (g *GrabBag) Float32(query string) float32 {
	res := g.Grab(query)
	if res != nil {
		return res.(float32)
	}
	return float32(0)
}

// Float64 grabs a query casted as a float64
func (g *GrabBag) Float64(query string) float64 {
	res := g.Grab(query)
	if res != nil {
		return res.(float64)
	}
	return float64(0)
}

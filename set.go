// Package set implements set data structure
package set

// Any type define an alias for the generic interface{} type
type Any interface{}

// Set struct implements set data structure for any type
type Set struct {
	elements map[Any]Any
}

// NewSet function creates a set and returns it
func NewSet() *Set {
	s := &Set{}
	s.elements = make(map[Any]Any)
	return s
}

// Add function add an item to the set
func (s *Set) Add(value Any) {
	s.elements[value] = nil
}

// AddSplittedString create a set entry foreach rune in the string
func (s *Set) AddSplittedString(value string) {
	for _, v := range value {
		s.elements[v] = nil
	}
}

// Remove function remove an items from the set
func (s *Set) Remove(value Any) {
	delete(s.elements, value)
}

// Contains function check if the set data structure contains the value element
func (s *Set) Contains(value Any) bool {
	_, c := s.elements[value]
	return c
}

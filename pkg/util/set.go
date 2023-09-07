package util

type Set struct {
	m map[interface{}]struct{}
}

// NewSet returns a new Set.
func NewSet() *Set {
	return &Set{
		m: make(map[interface{}]struct{}),
	}
}

// Add adds a new element to the set.
func (s *Set) Add(item interface{}) {
	s.m[item] = struct{}{}
}

// Remove removes an element from the set.
func (s *Set) Remove(item interface{}) {
	delete(s.m, item)
}

func (s *Set) Contains(item interface{}) bool {
	_, exists := s.m[item]
	return exists
}

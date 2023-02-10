package set

import "encoding/json"

// Set is a container that store unique elements in no particular order.
// Sets are an alias for map[T]struct{} where T is a comparable type.
type Set[T comparable] map[T]struct{}

// NewSet creates a new Set from a list of values.
func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T], len(values))
	set.Add(values...)

	return set
}

// Add adds a value to the set.
func (s Set[T]) Add(values ...T) {
	for _, value := range values {
		s[value] = struct{}{}
	}
}

// Contains returns true if the set contains the value.
func (s Set[T]) Contains(value T) bool {
	_, found := s[value]

	return found
}

// ContainsAll returns true if the set contains all the values.
func (s Set[T]) ContainsAll(values ...T) bool {
	for _, value := range values {
		if !s.Contains(value) {
			return false
		}
	}

	return true
}

// Remove removes a value from the set. Returns true if the set contained the value.
func (s Set[T]) Remove(value T) bool {
	if s.Contains(value) {
		delete(s, value)

		return true
	}

	return false
}

// RemoveAll removes all values from the set.
func (s Set[T]) RemoveAll(values ...T) {
	for _, value := range values {
		s.Remove(value)
	}
}

// Slice returns a slice of the values in the set.
func (s Set[T]) Slice() []T {
	slice := make([]T, 0, len(s))

	for k := range s {
		slice = append(slice, k)
	}

	return slice
}

func (s Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Slice())
}

func (s *Set[T]) UnmarshalJSON(data []byte) error {
	if (*s) == nil {
		*s = make(Set[T])
	}

	slice := make([]T, 0)
	err := json.Unmarshal(data, &slice)
	if err != nil {
		return err
	}

	s.Add(slice...)

	return nil
}

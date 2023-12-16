package collections

type Set[T comparable] map[T]struct{}


func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

func (s Set[T]) Add(element T) {
	s[element] = struct{}{}
}

func (s Set[T]) Remove(element T) {
	delete(s, element)
}

func (s Set[T]) Contains(element T) bool {
	_, exists := s[element]
	return exists
}

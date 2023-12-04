package collections

type StrSet map[string]struct{}

func (s StrSet) Add(element string) {
	s[element] = struct{}{}
}

func (s StrSet) Remove(element string) {
	delete(s, element)
}

func (s StrSet) Contains(element string) bool {
	_, exists := s[element]
	return exists
}

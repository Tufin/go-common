package common

type Set struct {
	// Using an empty struct{} has advantage that it doesn't require any additional space
	// Go's internal map type is optimized for that kind of values
	data map[string]struct{}
}

func NewSet() *Set {

	return &Set{data: make(map[string]struct{})}
}

func (s *Set) AddItems(items []string) *Set {

	for _, curr := range items {
		s.Add(curr)
	}

	return s
}

func (s *Set) Add(item string) *Set {

	s.data[item] = struct{}{}

	return s
}

func (s Set) Contains(item string) bool {

	_, ret := s.data[item]

	return ret
}

func (s Set) Size() int {

	return len(s.data)
}

func (s Set) Items() []string {

	if len(s.data) == 0 {
		return []string{}
	}

	ret := make([]string, 0, len(s.data))
	for i := range s.data {
		ret = append(ret, i)
	}

	return ret
}

// IsSimilar returns true if both set contains same items
func (s Set) IsSimilar(set []string) bool {

	ret := true
	if s.Size() != len(set) {
		ret = false
	} else {
		for _, currItem := range set {
			if !s.Contains(currItem) {
				ret = false
				break
			}
		}
	}

	return ret
}

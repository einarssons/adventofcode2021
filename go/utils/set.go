package utils

type Set map[string]struct{}

func CreateSet() Set {
	return Set(make(map[string]struct{}))
}

func (s Set) Contains(elem string) bool {
	_, ok := s[elem]
	return ok
}

func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

// Extend - extend set s with all elements in other (the result is union)
func (s Set) Extend(other Set) {
	for k := range other {
		s[k] = struct{}{}
	}
}

// Subtract - remove all elements from s that are in other
func (s Set) Subtract(other Set) {
	for k := range other {
		_, ok := s[k]
		if ok {
			delete(s, k)
		}
	}
}

// Intersect - only keep elements in s which are also in other
func (s Set) Intersect(other Set) {
	deleteList := make([]string, 0, len(s))
	for k := range s {
		_, inOther := other[k]
		if !inOther {
			deleteList = append(deleteList, k)
		}
	}
	for _, k := range deleteList {
		delete(s, k)
	}
}

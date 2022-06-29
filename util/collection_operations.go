package util

func FindFirst[T any](s []T, predicate func(i T) bool) *T {
	for _, item := range s {
		if predicate(item) {
			return &item
		}
	}
	return nil
}

// Creates the union of two slices. The item in the second slice will replace the one in the first slice
// if the hash value is the same. Items in the second slice which are not in the first will be added.
func Union[T any](s1 []T, s2 []T, hasher func(i *T) string) []T {
	result := []T{}

	m2 := ToMap(s2, hasher)

	// go through the first slice
	m1 := make(map[string]T, len(s1))
	for _, s1Item := range s1 {
		s1ItemHash := hasher(&s1Item)
		s2Item, exist := m2[s1ItemHash]
		if exist {
			// the item in the first slice will be replaces by the one of the second slice
			result = append(result, s2Item)
		} else {
			result = append(result, s1Item)
		}

		m1[s1ItemHash] = s1Item
	}

	// append items which were not in the first slice
	for _, s2Item := range s2 {
		s2ItemHash := hasher(&s2Item)
		_, exist := m1[s2ItemHash]
		if !exist {
			result = append(result, s2Item)
		}
	}

	return result
}

func ToMap[T any](s []T, hasher func(i *T) string) map[string]T {
	result := make(map[string]T, len(s))
	for _, item := range s {
		result[hasher(&item)] = item
	}
	return result
}

func Values[K comparable, V any](m map[K]V) []V {
	ret := make([]V, 0, len(m))
	for _, v := range m {
		ret = append(ret, v)
	}
	return ret
}
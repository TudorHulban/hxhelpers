package hxhelpers

func NotInSliceSource[T comparable](source []T, elements ...T) []T {
	if len(elements) == 0 {
		return nil
	}

	if len(source) == 0 {
		return elements // If source is empty, all elements are "not in source"
	}

	// Build a map of the source items for O(1) lookups
	sourceMap := make(map[T]struct{}, len(source))

	for _, item := range source {
		sourceMap[item] = struct{}{}
	}

	// Pre-allocate the result slice with a reasonable capacity estimate
	result := make([]T, 0, len(elements))

	// Filter the elements
	for _, element := range elements {
		if _, found := sourceMap[element]; !found {
			result = append(result, element)
		}
	}

	return result
}

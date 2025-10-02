package utils

func HasIntersection[T comparable](slice1, slice2 []T) bool {
	if len(slice1) == 0 || len(slice2) == 0 {
		return false
	}

	shorter, longer := slice1, slice2
	if len(slice2) < len(slice1) {
		shorter, longer = slice2, slice1
	}

	set := make(map[T]bool, len(shorter))
	for _, v := range shorter {
		set[v] = true
	}

	for _, v := range longer {
		if set[v] {
			return true
		}
	}
	return false
}

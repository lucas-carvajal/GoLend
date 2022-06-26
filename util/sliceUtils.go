package util

func Filter(elements []string, filter func(string) bool) []string {
	var res []string
	for _, elem := range elements {
		if filter(elem) {
			res = append(res, elem)
		}
	}
	return res
}

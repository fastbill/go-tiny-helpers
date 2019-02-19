// Package maputil contains helpers for working with maps.
package maputil

// Keys returns the keys of a map.
// The keys and values of the map must be strings.
func Keys(input map[string]string) []string {
	result := make([]string, len(input))
	i := 0
	for k := range input {
		result[i] = k
		i++
	}

	return result
}

// Values returns the values of a map.
// The keys and values must be strings.
func Values(input map[string]string) []string {
	result := make([]string, len(input))
	i := 0
	for _, v := range input {
		result[i] = v
		i++
	}

	return result
}

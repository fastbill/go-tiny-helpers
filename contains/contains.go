// Package contains checks if a value is included in a slice of values.
package contains

// String checks if a slice of strings contains a certain string value.
func String(list []string, value string) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// Int checks if a slice of integers contains a certain int value.
func Int(list []int, value int) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// Int64 checks if a slice of int64 values contains a certain int64 value.
func Int64(list []int64, value int64) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

// Uint64 checks if a slice of uint64 values contains a certain uint64 value.
func Uint64(list []uint64, value uint64) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

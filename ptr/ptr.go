// Package ptr contains helpers to convert values into pointers to these values.
// This is useful when working with structs that are used to parse JSON input. Pointer types are often used there to account for fields being null or absent.
// To fill these structs with dummy values for tests you need to create pointers to values with primitive types (see example).
package ptr

// String returns a pointer to the given string value.
func String(v string) *string {
	return &v
}

// Int returns a pointer to a given integer value.
func Int(v int) *int {
	return &v
}

// Int64 returns a pointer to a given int64 value.
func Int64(v int64) *int64 {
	return &v
}

// Uint64 returns a pointer to a given int64 value.
func Uint64(v uint64) *uint64 {
	return &v
}

// Float64 returns a pointer to a given float64 value.
func Float64(v float64) *float64 {
	return &v
}

// Bool returns a pointer to a given bool value.
func Bool(v bool) *bool {
	return &v
}

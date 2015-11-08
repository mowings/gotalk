// Package stuff provides some examples for demonstrating go's test
// facilities
package stuff

// Reverse a string. It's pretty neat
// This function takes a string and returns a string
func Reverse(s string) string {
	o := make([]byte, len(s))
	i := len(o)
	for j := 0; j < len(o); j++ {
		i--
		o[i] = s[j]
	}
	return string(o)
}

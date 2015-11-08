package stuff

func Reverse(s string) string {
	o := make([]byte, len(s))
	i := len(o)
	for j := 0; j < len(o); j++ {
		i--
		o[i] = s[j]
	}
	return string(o)
}

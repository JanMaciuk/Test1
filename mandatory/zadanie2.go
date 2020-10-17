package mandatory

//Concatenate returns a slice
func Concatenate(s1, s2 []int) []int {
	var a = append(s1, s2...)
	return a
}

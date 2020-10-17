package mandatory

//FilterEven returns odd numbers
func FilterEven(age []int) []int {
	var slajs []int
	for _, i := range age {
		if i%2 != 0 {
			slajs = append(slajs, i)
		}
	}
	return slajs
}

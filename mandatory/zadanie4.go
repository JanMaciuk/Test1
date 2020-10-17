package mandatory

//PrintDividedBy4 prints all numbers divisible by 4 from 1 to given n
func PrintDividedBy4(n int) []int {
	var slajs []int
	for i := 1; i <= n; i++ {
		if i%4 == 0 {
			slajs = append(slajs, i)
		}
	}
	return slajs
}

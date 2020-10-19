package optional

import "flag"

var slowo = flag.String("napis", "jakub", "given string")

//StringLengthFlag returns the lenght of the given string
func StringLengthFlag(slowo string) int {

	return len(slowo)
}

package main

import (
	"fmt"
	"github/jahopl/Test1/mandatory"
	"github/jahopl/Test1/optional"
)

func main() {
	u := mandatory.User{
		Name:     "andrzejek",
		LastName: "kowalszewski",
		Age:      12,
		IsAdult:  false,
	}

	//	fmt.Print(mandatory.Suma(1, 2))
	//	fmt.Print(mandatory.Concatenate([]int{1, 2}, []int{3, 4}))
	//	fmt.Print(mandatory.Is18(18))
	//	fmt.Print(mandatory.PrintDividedBy4(120))
	//	fmt.Print(mandatory.FilterEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	//	fmt.Print(mandatory.NewUser("Andrzej", "Kowalski", 21))
	u.ChangeAge(15)
	fmt.Print(u.ToString())
	fmt.Print(optional.StringLength("aaa"))

}

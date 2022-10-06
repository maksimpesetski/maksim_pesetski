package main

import (
	"fmt"

	"github.com/maksimpesetski/maksim_pesetski/internal/mapper"
	"github.com/maksimpesetski/maksim_pesetski/internal/problems"
)

/*
 	And change your code to look like this:
	func main() {
   		s := NewSkipString(3, "Aspiration.com")
   		mapper.MapString(&s)
   		fmt.Println(s)
	}
	Make sure your fmt.Println(s) output looks nice and clean, ie implement the Stringer interface.
*/

func main() {

	input := "Aspiration.com"
	/*
		NOTE:
			I could keep aspiration-mapper-problem-part-2.go at the same level as main.go in order to invoke
			"s := problems.NewSkipString(3, input)" like "s := NewSkipString(3, input)", but I think it's cleaner to keep
			the solution in problems package
	*/
	s := problems.NewSkipString(3, input)
	mapper.MapString(&s)
	/*
		NOTE:
			We could print the result in much cleaner way by using fmt.Printf("Processed '%s' sting into '%s'", input, s),
			but I assume you specifically wanted me to use fmt.Println(s).
	*/
	fmt.Println("Processed ", input, " string into ", s, "")
}

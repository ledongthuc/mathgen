package mathgen_test

import (
	"fmt"

	"github.com/ledongthuc/mathgen"
)

func Example_Add_two_numbers() {
	questionAndAnswers, _ := mathgen.AddIntegers(20)
	fmt.Println("Addends:", questionAndAnswers.Addends)
	fmt.Println("Sum:", questionAndAnswers.Sum)
	fmt.Println("Format:", questionAndAnswers.String())
}

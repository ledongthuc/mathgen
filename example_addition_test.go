package mathgen_test

import (
	"fmt"

	"github.com/ledongthuc/mathgen"
)

func ExampleAddIntegers() {
	questionAndAnswers, _ := mathgen.AddIntegers(20)
	fmt.Println("Addends:", questionAndAnswers.Addends)
	fmt.Println("Sum:", questionAndAnswers.Sum)
	fmt.Println("Print Question:", questionAndAnswers.StringQuestion())
	fmt.Println("Print Full:", questionAndAnswers.String())
}

func ExampleAddIntegerN() {
	questionAndAnswers, _ := mathgen.AddIntegerN(3, 20)
	fmt.Println("Addends:", questionAndAnswers.Addends)
	fmt.Println("Sum:", questionAndAnswers.Sum)
	fmt.Println("Print Question:", questionAndAnswers.StringQuestion())
	fmt.Println("Print Full:", questionAndAnswers.String())
}

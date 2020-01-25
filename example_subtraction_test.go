package mathgen_test

import (
	"fmt"

	"github.com/ledongthuc/mathgen"
)

func ExampleSubtractIntegers() {
	questionAndAnswers, _ := mathgen.SubtractIntegers(20)
	fmt.Println("Minuend:", questionAndAnswers.Minuend)
	fmt.Println("Subtrahends:", questionAndAnswers.Subtrahends)
	fmt.Println("Difference:", questionAndAnswers.Difference)
	fmt.Println("Print Question:", questionAndAnswers.StringQuestion())
	fmt.Println("Print Full:", questionAndAnswers.String())
}

func ExampleSubtractIntegerN() {
	questionAndAnswers, _ := mathgen.SubtractIntegerN(3, 20)
	fmt.Println("Minuend:", questionAndAnswers.Minuend)
	fmt.Println("Subtrahends:", questionAndAnswers.Subtrahends)
	fmt.Println("Difference:", questionAndAnswers.Difference)
	fmt.Println("Print Question:", questionAndAnswers.StringQuestion())
	fmt.Println("Print Full:", questionAndAnswers.String())
}

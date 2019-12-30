package mathgen_test

import (
	"fmt"

	"github.com/ledongthuc/mathgen"
)

func ExampleAddIntegers() {
	questionAndAnswers, _ := mathgen.AddIntegers(20)
	fmt.Println("Addends:", questionAndAnswers.Addends)
	fmt.Println("Sum:", questionAndAnswers.Sum)
	fmt.Println("Format:", questionAndAnswers.String())

	// Output:
	// Addends: [9 8]
	// Sum: 17
	// Format: 9 + 8 = 17
}

func ExampleAddIntegerN() {
	questionAndAnswers, _ := mathgen.AddIntegerN(3, 20)
	fmt.Println("Addends:", questionAndAnswers.Addends)
	fmt.Println("Sum:", questionAndAnswers.Sum)
	fmt.Println("Format:", questionAndAnswers.String())

	// Output:
	// Addends: [9 8]
	// Sum: 17
	// Format: 9 + 8 = 17
}

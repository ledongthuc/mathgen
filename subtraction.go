package mathgen

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// SubtractionResult of addition operator
type SubtractionResult struct {
	Minuend     int64
	Subtrahends []int64
	Difference  int64
}

// String of SubtractionResult will present struct data with question answer format a - b - c - d ... = difference
func (s SubtractionResult) String() string {
	return s.StringQuestion() + strconv.FormatInt(s.Difference, 10)
}

// StringQuestion of SubtractionResult will present struct data with question format a - b - c - d ... =
func (s SubtractionResult) StringQuestion() string {
	if len(s.Subtrahends) == 0 {
		return fmt.Sprintf("%d = ", s.Minuend)
	}

	sep := " - "
	var b strings.Builder
	b.Grow((len(sep) + 5) * (len(s.Subtrahends) - 1))
	b.WriteString(strconv.FormatInt(s.Minuend, 10))
	for _, subtrahend := range s.Subtrahends {
		b.WriteString(sep)
		b.WriteString(strconv.FormatInt(subtrahend, 10))
	}
	b.WriteString(" = ")
	return b.String()
}

// SubtractIntegers generate a operator with subtraction of 1 minuend and 1 subtrahend.
// Parameter maxMinuend defines maxsimum of minuend, it should be greater than 0.
// If maxMinuend = 1, addends always is 1 - 1
func SubtractIntegers(maxMinuend int64) (SubtractionResult, error) {
	return subtractIntegerN(getRand(), 2, maxMinuend)
}

// SubtractIntegerN generate a operator with subtraction.
// Parameter numberOfSubtrahends defines subtrahends in question, it should be greater than 0.
// Parameter maxMinuend defines maxsimum of minuend result, it should be greater than 0.
// If maxMinuend = 1, addends always is 1 - 1 - 0 ....
func SubtractIntegerN(numberOfSubtrahends int, maxMinuend int64) (SubtractionResult, error) {
	return subtractIntegerN(getRand(), numberOfSubtrahends, maxMinuend)
}

// subtractIntegerN generate a operator with subtraction.
// r is source of random numbers
// Parameter numberOfSubtrahends defines subtrahends in question, it should be greater than 0.
// Parameter maxMinuend defines maxsimum of minuend result, it should be greater than 0.
// If maxMinuend = 1, addends always is 1 - 1 - 0 ....
func subtractIntegerN(r *rand.Rand, numberOfSubtrahends int, maxMinuend int64) (SubtractionResult, error) {
	if numberOfSubtrahends <= 0 {
		return SubtractionResult{}, errors.New("number of subtrahend should be greater than 0")
	}
	if maxMinuend <= 0 {
		return SubtractionResult{}, errors.New("max minuend should be greater than 0")
	}
	if maxMinuend <= int64(numberOfSubtrahends) {
		subtrahends := make([]int64, maxMinuend, maxMinuend)
		for i := int64(0); i < maxMinuend; i++ {
			subtrahends[i] = 1
		}
		return SubtractionResult{Minuend: maxMinuend, Subtrahends: subtrahends, Difference: 0}, nil
	}

	additionResult, err := addIntegerN(r, numberOfSubtrahends+1, maxMinuend)
	if err != nil {
		return SubtractionResult{}, err
	}
	return SubtractionResult{
		Minuend:     additionResult.Sum,
		Subtrahends: additionResult.Addends[:len(additionResult.Addends)-1],
		Difference:  additionResult.Addends[len(additionResult.Addends)-1],
	}, nil
}

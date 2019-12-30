package mathgen

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// Addition Result of addition operator
type AdditionResult struct {
	Addends []int64
	Sum     int64
}

// String of AdditionResult will present struct data with question answer format a + b + c + ... = sum
func (a AdditionResult) String() string {
	switch len(a.Addends) {
	case 0:
		return strconv.FormatInt(a.Sum, 10)
	case 1:
		return fmt.Sprintf("%d = %d", a.Addends[0], a.Sum)
	}

	sep := " + "
	var b strings.Builder
	b.Grow((len(sep) + 5) * (len(a.Addends) - 1))
	b.WriteString(strconv.FormatInt(a.Addends[0], 10))
	for _, addend := range a.Addends[1:] {
		b.WriteString(sep)
		b.WriteString(strconv.FormatInt(addend, 10))
	}
	b.WriteString(" = ")
	b.WriteString(strconv.FormatInt(a.Sum, 10))
	return b.String()
}

// AddIntegers generate a operator with sum of 2 addends.
// Parameter maxSum defines maxsimum of Sum result, it should be greater than 0.
// If maxSum = 2, addends always is 1 + 1
func AddIntegers(maxSum int64) (AdditionResult, error) {
	return addIntegerN(getRand(), 2, maxSum)
}

// AddIntegerN generate a operator with sum.
// Parameter numberOfAddend defines addends in question, it should be greater than 1.
// Parameter maxSum defines maxsimum of Sum result, it should be greater than 0.
// If maxSum is greater than or equal numberOfAddend. Question always has pattern: 1 + 1 + 1 ....
func AddIntegerN(numberOfAddend int, maxSum int64) (AdditionResult, error) {
	return addIntegerN(getRand(), numberOfAddend, maxSum)
}

// addIntegerN generate a operator with sum.
// r is  source of random numbers
// Parameter numberOfAddend defines addends in question, it should be greater than 1.
// Parameter maxSum defines maxsimum of Sum result, it should be greater than 0.
// If maxSum is greater than or equal numberOfAddend. Question always has pattern: 1 + 1 + 1 ....
func addIntegerN(r *rand.Rand, numberOfAddend int, maxSum int64) (AdditionResult, error) {
	if numberOfAddend <= 1 {
		return AdditionResult{}, errors.New("number of addend should be greater than 1")
	}
	if maxSum <= 1 {
		return AdditionResult{}, errors.New("max sum should be greater than 0")
	}
	if maxSum <= int64(numberOfAddend) {
		addends := make([]int64, maxSum, maxSum)
		for i := int64(0); i < maxSum; i++ {
			addends[i] = 1
		}
		return AdditionResult{Addends: addends, Sum: maxSum}, nil
	}

	// Generate fake addend and sum that will use to calculate percent of real SUM
	result := AdditionResult{Addends: make([]int64, 0, numberOfAddend)}
	for i := 0; i < numberOfAddend; i++ {
		// make sure the fake addend doesn't reach to limit int64
		// random from [0-9) then +1 and we will have [1-10)
		addend := int64(rand.Int31n(9) + 1)
		result.Addends = append(result.Addends, addend)
		result.Sum += addend // Fake sum
	}

	// Real sum will be generated with (numberOfAddend:SumMax]
	realSum := r.Int63n(maxSum+1-int64(numberOfAddend)) + int64(numberOfAddend)
	var calculatingSum int64
	for index, fakeAddend := range result.Addends {
		var addend int64
		if index == len(result.Addends)-1 {
			addend = realSum - calculatingSum
		} else {
			addend = int64(math.Round(float64(fakeAddend) / float64(result.Sum) * float64(realSum)))
			calculatingSum += addend
		}
		result.Addends[index] = addend
	}
	result.Sum = realSum
	return result, nil
}

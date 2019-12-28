package mathgen

import (
	"math"
	"math/rand"
)

type AdditionResult struct {
	Addends []int64
	Sum     int64
}

func AddIntegers(maxSum int64) AdditionResult {
	return addIntegerN(getRand(), 2, maxSum)
}

func AddIntegerN(numberOfAddend int, maxSum int64) AdditionResult {
	return addIntegerN(getRand(), numberOfAddend, maxSum)
}

func addIntegerN(r *rand.Rand, numberOfAddend int, maxSum int64) AdditionResult {
	if numberOfAddend == 0 || maxSum == 0 {
		return AdditionResult{}
	}
	if maxSum <= int64(numberOfAddend) {
		addends := make([]int64, maxSum, maxSum)
		for i := int64(0); i < maxSum; i++ {
			addends[i] = 1
		}
		return AdditionResult{Addends: addends, Sum: maxSum}
	}

	// Generate fake addend and sum that will use to calculate percent of real SUM
	result := AdditionResult{Addends: make([]int64, 0, numberOfAddend)}
	for i := 0; i < numberOfAddend; i++ {
		addend := int64(rand.Int31()) // make sure the fake addend doesn't reach to limit int64
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
	return result
}

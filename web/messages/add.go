package messages

import (
	"fmt"

	"github.com/ledongthuc/mathgen"
)

type AdditionResult struct {
	Addends []int64 `json:"addends"`
	Sum     int64   `json:"sum"`
}

func AdditionResultFromModel(model mathgen.AdditionResult) AdditionResult {
	result := AdditionResult{
		Sum:     model.Sum,
		Addends: make([]int64, len(model.Addends)),
	}
	copy(result.Addends, model.Addends)
	return result
}

type AdditionRequest struct {
	NumberOfAddend int   `json:"number_of_addend"`
	MaxSum         int64 `json:"max_sum"`
}

func (a AdditionRequest) Valid() (bool, error) {
	if a.NumberOfAddend == 0 {
		return false, fmt.Errorf("Number of addend's required")
	}
	if a.MaxSum == 0 {
		return false, fmt.Errorf("Maximum sum should be greater than 0")
	}
	if a.MaxSum < int64(a.NumberOfAddend) {
		return false, fmt.Errorf("Maximum sum should be greater than or equal %d", int64(a.NumberOfAddend))
	}
	return true, nil
}

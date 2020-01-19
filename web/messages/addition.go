package messages

import (
	"fmt"

	"github.com/ledongthuc/mathgen"
)

type AdditionResult struct {
	Addends  []int64 `json:"addends"`
	Sum      int64   `json:"sum"`
	Question string  `json:"question"`
	Result   string  `json:"result"`
}

func AdditionResultFromModel(model mathgen.AdditionResult) AdditionResult {
	result := AdditionResult{
		Sum:      model.Sum,
		Addends:  make([]int64, len(model.Addends)),
		Question: model.StringQuestion(),
		Result:   model.String(),
	}
	copy(result.Addends, model.Addends)
	return result
}

type AdditionRequest struct {
	NumberOfAddends int   `json:"number_of_addends"`
	MaxSum          int64 `json:"max_sum"`
}

func (a AdditionRequest) Valid() (bool, error) {
	if a.NumberOfAddends == 0 {
		return false, fmt.Errorf("Number of addends's required")
	}
	if a.MaxSum == 0 {
		return false, fmt.Errorf("Maximum sum should be greater than 0")
	}
	if a.MaxSum < int64(a.NumberOfAddends) {
		return false, fmt.Errorf("Maximum sum should be greater than or equal %d", int64(a.NumberOfAddends))
	}
	return true, nil
}

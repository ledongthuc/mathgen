package messages

import (
	"fmt"

	"github.com/ledongthuc/mathgen"
)

// AdditionResult is result message that will be used to return on API
type AdditionResult struct {
	Addends  []int64 `json:"addends"`
	Sum      int64   `json:"sum"`
	Question string  `json:"question"`
	Result   string  `json:"result"`
}

// AdditionResultFromModel converts message from web to request model of Addition
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

// AdditionRequest is request to make a add generation from API
type AdditionRequest struct {
	NumberOfAddends int   `json:"number_of_addends"`
	MaxSum          int64 `json:"max_sum"`
}

// Valid of AdditionRequest check the validation of the request message
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

package messages

import (
	"fmt"

	"github.com/ledongthuc/mathgen"
)

// SubtractionResult is result message that will be used to return on API
type SubtractionResult struct {
	Minuend     int64   `json:"minuend"`
	Subtrahends []int64 `json:"subtrahends"`
	Difference  int64   `json:"difference"`
	Question    string  `json:"question"`
	Result      string  `json:"result"`
}

// SubtractionResultFromModel converts message from web to request model of Subtraction
func SubtractionResultFromModel(model mathgen.SubtractionResult) SubtractionResult {
	result := SubtractionResult{
		Minuend:     model.Minuend,
		Subtrahends: make([]int64, len(model.Subtrahends)),
		Difference:  model.Difference,
		Question:    model.StringQuestion(),
		Result:      model.String(),
	}
	copy(result.Subtrahends, model.Subtrahends)
	return result
}

// SubtractionRequest is request to make a subtract generation from API
type SubtractionRequest struct {
	MaxMinuend          int64 `json:"max_minuend"`
	NumberOfSubtrahends int   `json:"number_of_subtrahends"`
}

// Valid of SubtractionRequest check the validation of the request message
func (a SubtractionRequest) Valid() (bool, error) {
	if a.NumberOfSubtrahends == 0 {
		return false, fmt.Errorf("Number of subtrahend's required")
	}
	if a.MaxMinuend == 0 {
		return false, fmt.Errorf("Maximum minuend should be greater than 0")
	}
	if a.MaxMinuend < int64(a.NumberOfSubtrahends) {
		return false, fmt.Errorf("Maximum subtrahend should be greater than or equal %d", int64(a.NumberOfSubtrahends))
	}
	return true, nil
}

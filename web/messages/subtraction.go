package messages

import (
	"fmt"

	"github.com/ledongthuc/mathgen"
)

type SubtractionResult struct {
	Minuend     int64   `json:"minuend"`
	Subtrahends []int64 `json:"subtrahends"`
	Difference  int64   `json:"difference"`
	Question    string  `json:"question"`
	Result      string  `json:"result"`
}

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

type SubtractionRequest struct {
	MaxMinuend          int64 `json:"max_minuend"`
	NumberOfSubtrahends int   `json:"number_of_subtrahends"`
}

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

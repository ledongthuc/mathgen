package mathgen

import (
	"fmt"
	"strconv"
	"strings"
)

// Substraction Result of addition operator
type SubstractionResult struct {
	Minuend     int64
	Subtrahends []int64
	Difference  int64
}

// String of SubstractionResult will present struct data with question answer format a - b - c - d ... = difference
func (s SubstractionResult) String() string {
	return s.StringQuestion() + strconv.FormatInt(s.Difference, 10)
}

// StringQuestion of SubstractionResult will present struct data with question format a - b - c - d ... =
func (s SubstractionResult) StringQuestion() string {
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

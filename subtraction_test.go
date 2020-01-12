package mathgen

import (
	"testing"
)

func TestSubstractionResult_StringQuestion(t *testing.T) {
	type fields struct {
		Minuend     int64
		Subtrahends []int64
		Difference  int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Lengh of subtrahends is 0",
			fields: fields{
				Minuend:     10,
				Subtrahends: []int64{},
				Difference:  10,
			},
			want: "10 = ",
		},
		{
			name: "Lengh of subtrahends is 1",
			fields: fields{
				Minuend:     10,
				Subtrahends: []int64{5},
				Difference:  5,
			},
			want: "10 - 5 = ",
		},
		{
			name: "Lengh of subtrahends is 3",
			fields: fields{
				Minuend:     10,
				Subtrahends: []int64{1, 2, 3},
				Difference:  4,
			},
			want: "10 - 1 - 2 - 3 = ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubstractionResult{
				Minuend:     tt.fields.Minuend,
				Subtrahends: tt.fields.Subtrahends,
				Difference:  tt.fields.Difference,
			}
			if got := s.StringQuestion(); got != tt.want {
				t.Errorf("SubstractionResult.StringQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstractionResult_String(t *testing.T) {
	type fields struct {
		Minuend     int64
		Subtrahends []int64
		Difference  int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Lengh of subtrahends is 0",
			fields: fields{
				Minuend:     10,
				Subtrahends: []int64{},
				Difference:  10,
			},
			want: "10 = 10",
		},
		{
			name: "Lengh of subtrahends is 1",
			fields: fields{
				Minuend:     10,
				Subtrahends: []int64{5},
				Difference:  5,
			},
			want: "10 - 5 = 5",
		},
		{
			name: "Lengh of subtrahends is 3",
			fields: fields{
				Minuend:     10,
				Subtrahends: []int64{1, 2, 3},
				Difference:  4,
			},
			want: "10 - 1 - 2 - 3 = 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SubstractionResult{
				Minuend:     tt.fields.Minuend,
				Subtrahends: tt.fields.Subtrahends,
				Difference:  tt.fields.Difference,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("SubstractionResult.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

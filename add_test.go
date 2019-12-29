package mathgen

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_addIntegerN(t *testing.T) {
	type args struct {
		r              *rand.Rand
		numberOfAddend int
		sumMax         int64
	}
	tests := []struct {
		name string
		args args
		want AdditionResult
	}{
		{
			name: "Don't have any addend",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 0,
				sumMax:         10,
			},
			want: AdditionResult{
				Addends: nil,
				Sum:     0,
			},
		},
		{
			name: "Max sum is 0",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 3,
				sumMax:         0,
			},
			want: AdditionResult{
				Addends: nil,
				Sum:     0,
			},
		},
		{
			name: "Number of addend is equal with sum max",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 5,
				sumMax:         5,
			},
			want: AdditionResult{
				Addends: []int64{1, 1, 1, 1, 1},
				Sum:     5,
			},
		},
		{
			name: "Sum max 5 - 1st",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 2,
				sumMax:         5,
			},
			want: AdditionResult{
				Addends: []int64{1, 2},
				Sum:     3,
			},
		},
		{
			name: "Sum max 5 - 2nd",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 2,
				sumMax:         5,
			},
			want: AdditionResult{
				Addends: []int64{2, 1},
				Sum:     3,
			},
		},
		{
			name: "Successful - large number 1st",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 3,
				sumMax:         1000,
			},
			want: AdditionResult{
				Addends: []int64{133, 186, 213},
				Sum:     532,
			},
		},
		{
			name: "Successful - large number 2st",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 3,
				sumMax:         1000,
			},
			want: AdditionResult{
				Addends: []int64{228, 127, 177},
				Sum:     532,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addIntegerN(tt.args.r, tt.args.numberOfAddend, tt.args.sumMax); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addIntegerN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addIntegerMass(t *testing.T) {
	type args struct {
		r              *rand.Rand
		numberOfAddend int
		sumMax         int64
	}
	tt := struct {
		name string
		args args
	}{
		name: "Successful",
		args: args{
			r:              rand.New(rand.NewSource(0)),
			numberOfAddend: 100,
			sumMax:         1000,
		},
	}
	for i := 0; i < 10000; i++ {
		t.Run(tt.name, func(t *testing.T) {
			result := addIntegerN(tt.args.r, tt.args.numberOfAddend, tt.args.sumMax)
			var actualSum int64
			for _, addend := range result.Addends {
				actualSum += addend
			}
			if actualSum != result.Sum {
				t.Errorf("addIntegerN() at index %v = %v, want %v", i, actualSum, result.Sum)
			}
		})
	}
}

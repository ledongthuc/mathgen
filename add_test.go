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
		name    string
		args    args
		want    AdditionResult
		wantErr bool
	}{
		{
			name: "Don't have any addend",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 0,
				sumMax:         10,
			},
			want:    AdditionResult{},
			wantErr: true,
		},
		{
			name: "Only one addend",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 1,
				sumMax:         10,
			},
			want:    AdditionResult{},
			wantErr: true,
		},
		{
			name: "Max sum is 0",
			args: args{
				r:              rand.New(rand.NewSource(0)),
				numberOfAddend: 3,
				sumMax:         0,
			},
			want:    AdditionResult{},
			wantErr: true,
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
			wantErr: false,
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
			wantErr: false,
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
			wantErr: false,
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
			wantErr: false,
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addIntegerN(tt.args.r, tt.args.numberOfAddend, tt.args.sumMax)
			if (err != nil) != tt.wantErr {
				t.Errorf("addIntegerN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
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
			result, _ := addIntegerN(tt.args.r, tt.args.numberOfAddend, tt.args.sumMax)
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

func TestAdditionResult_String(t *testing.T) {
	type fields struct {
		Addends []int64
		Sum     int64
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Lengh of addends is 0",
			fields: fields{
				Addends: []int64{},
				Sum:     0,
			},
			want: "0",
		},
		{
			name: "Lengh of addends is 1",
			fields: fields{
				Addends: []int64{5},
				Sum:     5,
			},
			want: "5 = 5",
		},
		{
			name: "Lengh of addends is 2",
			fields: fields{
				Addends: []int64{101, 67},
				Sum:     168,
			},
			want: "101 + 67 = 168",
		},
		{
			name: "Lengh of addends is 3",
			fields: fields{
				Addends: []int64{101, 67, 201},
				Sum:     369,
			},
			want: "101 + 67 + 201 = 369",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := AdditionResult{
				Addends: tt.fields.Addends,
				Sum:     tt.fields.Sum,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("AdditionResult.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

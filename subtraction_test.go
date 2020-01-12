package mathgen

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestSubtractionResult_StringQuestion(t *testing.T) {
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
			s := SubtractionResult{
				Minuend:     tt.fields.Minuend,
				Subtrahends: tt.fields.Subtrahends,
				Difference:  tt.fields.Difference,
			}
			if got := s.StringQuestion(); got != tt.want {
				t.Errorf("SubtractionResult.StringQuestion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtractionResult_String(t *testing.T) {
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
			s := SubtractionResult{
				Minuend:     tt.fields.Minuend,
				Subtrahends: tt.fields.Subtrahends,
				Difference:  tt.fields.Difference,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("SubtractionResult.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtractIntegerN(t *testing.T) {
	type args struct {
		r                   *rand.Rand
		numberOfSubtrahends int
		maxMinuend          int64
	}
	tests := []struct {
		name    string
		args    args
		want    SubtractionResult
		wantErr bool
	}{
		{
			name: "number of subtrahends is 0",
			args: args{
				r:                   rand.New(rand.NewSource(0)),
				numberOfSubtrahends: 0,
				maxMinuend:          10,
			},
			want:    SubtractionResult{},
			wantErr: true,
		},
		{
			name: "max minuend is 0",
			args: args{
				r:                   rand.New(rand.NewSource(0)),
				numberOfSubtrahends: 1,
				maxMinuend:          0,
			},
			want:    SubtractionResult{},
			wantErr: true,
		},
		{
			name: "max minuend equals number of subtrahends",
			args: args{
				r:                   rand.New(rand.NewSource(0)),
				numberOfSubtrahends: 5,
				maxMinuend:          5,
			},
			want: SubtractionResult{
				Minuend:     5,
				Subtrahends: []int64{1, 1, 1, 1, 1},
				Difference:  0,
			},
			wantErr: false,
		},
		{
			name: "max minuend is less than number of subtrahends",
			args: args{
				r:                   rand.New(rand.NewSource(0)),
				numberOfSubtrahends: 5,
				maxMinuend:          3,
			},
			want: SubtractionResult{
				Minuend:     3,
				Subtrahends: []int64{1, 1, 1},
				Difference:  0,
			},
			wantErr: false,
		},
		{
			name: "max minuend 5 - 1st",
			args: args{
				r:                   rand.New(rand.NewSource(0)),
				numberOfSubtrahends: 1,
				maxMinuend:          5,
			},
			want: SubtractionResult{
				Minuend:     3,
				Subtrahends: []int64{2},
				Difference:  1,
			},
			wantErr: false,
		},
		{
			name: "max minuend 5 - 2st",
			args: args{
				r:                   rand.New(rand.NewSource(0)),
				numberOfSubtrahends: 1,
				maxMinuend:          5,
			},
			want: SubtractionResult{
				Minuend:     3,
				Subtrahends: []int64{1},
				Difference:  2,
			},
			wantErr: false,
		},
		{
			name: "Successful - large number 1st",
			args: args{
				r:                   rand.New(rand.NewSource(0)),
				numberOfSubtrahends: 3,
				maxMinuend:          1000,
			},
			want: SubtractionResult{
				Minuend:     844,
				Subtrahends: []int64{270, 169, 169},
				Difference:  236,
			},
			wantErr: false,
		},
		{
			name: "Successful - large number 2nd",
			args: args{
				r:                   rand.New(rand.NewSource(0)),
				numberOfSubtrahends: 3,
				maxMinuend:          1000,
			},
			want: SubtractionResult{
				Minuend:     844,
				Subtrahends: []int64{106, 281, 141},
				Difference:  316,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subtractIntegerN(tt.args.r, tt.args.numberOfSubtrahends, tt.args.maxMinuend)
			if (err != nil) != tt.wantErr {
				t.Errorf("subtractIntegerN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtractIntegerN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtractIntegerMass(t *testing.T) {
	type args struct {
		r                   *rand.Rand
		numberOfSubtrahends int
		maxMinuend          int64
	}
	tt := struct {
		name string
		args args
	}{
		name: "Successful",
		args: args{
			r:                   rand.New(rand.NewSource(0)),
			numberOfSubtrahends: 100,
			maxMinuend:          1000,
		},
	}
	for i := 0; i < 10000; i++ {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subtractIntegerN(tt.args.r, tt.args.numberOfSubtrahends, tt.args.maxMinuend)
			if err != nil {
				t.Errorf("subtractIntegerN() error = %v", err)
			}
			actuallyDifference := got.Minuend
			for _, subtrahend := range got.Subtrahends {
				actuallyDifference -= subtrahend
			}
			if actuallyDifference != got.Difference {
				t.Errorf("subtractIntegerN() at index %v = %v, want %v", i, actuallyDifference, got.Difference)
			}
		})
	}
}

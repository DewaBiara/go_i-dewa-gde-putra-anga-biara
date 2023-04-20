package calculate

import (
	"testing"
)

func Test_Addition(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test addition int values",
			args: args{
				a: 1,
				b: 2,
			},
			want: 3,
		},
		{
			name: "Test addition float and int values",
			args: args{
				a: 1.5,
				b: 2,
			},
			want: 3.5,
		},
		{
			name: "Test addition float values",
			args: args{
				a: 1.5,
				b: 2.5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Subtraction(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test subtraction int values",
			args: args{
				a: 1,
				b: 2,
			},
			want: -1,
		},
		{
			name: "Test subtraction float and int values",
			args: args{
				a: 1.5,
				b: 2,
			},
			want: -0.5,
		},
		{
			name: "Test subtraction float values",
			args: args{
				a: 1.5,
				b: 2.5,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtraction(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("subtraction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Division(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test division int values",
			args: args{
				a: 1,
				b: 2,
			},
			want: 0.5,
		},
		{
			name: "Test division float and int values",
			args: args{
				a: 1.5,
				b: 2,
			},
			want: 0.75,
		},
		{
			name: "Test division float values",
			args: args{
				a: 1.5,
				b: 2.5,
			},
			want: 0.6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Division(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("division() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Multiplication(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test multiplication int values",
			args: args{
				a: 1,
				b: 2,
			},
			want: 2,
		},
		{
			name: "Test multiplication float and int values",
			args: args{
				a: 1.5,
				b: 2,
			},
			want: 3,
		},
		{
			name: "Test multiplication float values",
			args: args{
				a: 1.5,
				b: 2.5,
			},
			want: 3.75,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiplication(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("multiplication() = %v, want %v", got, tt.want)
			}
		})
	}
}

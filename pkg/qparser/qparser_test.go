package qparser

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want Result
	}{
		{
			name: "empty",
			args: args{
				args: []string{},
			},
			want: Result{
				Namespace: "",
				Query:     "",
			},
		},
		{
			name: "one",
			args: args{
				args: []string{"foo"},
			},
			want: Result{
				Namespace: "foo",
				Query:     "",
			},
		},
		{
			name: "two",
			args: args{
				args: []string{"foo", "bar"},
			},
			want: Result{
				Namespace: "foo",
				Query:     "bar",
			},
		},
		{
			name: "three",
			args: args{
				args: []string{"foo", "bar", "baz"},
			},
			want: Result{
				Namespace: "foo",
				Query:     "bar+baz",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

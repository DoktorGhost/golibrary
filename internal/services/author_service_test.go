package services

import "testing"

func Test_validName(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{name: "1",
			args: "Oleg",
			want: true,
		},
		{name: "2",
			args: "Oleg2",
			want: false},
		{name: "3",
			args: "",
			want: true},
		{name: "4",
			args: "32423",
			want: false},
		{name: "5",
			args: "Салтыков-Щедрин",
			want: true},
		{name: "6",
			args: "-Салтыков-Щедрин",
			want: false},
		{name: "7",
			args: "Салтыков-Щедрин-",
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validName(tt.args); got != tt.want {
				t.Errorf("validName() = %v, want %v", got, tt.want)
			}
		})
	}
}

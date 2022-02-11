package formattingexample

import (
	"fmt"
	"testing"
)

func TestInitFormattingExample(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				x: 1,
				y: 2,
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitFormattingExample(tt.args.x, tt.args.y)
			fmt.Println((err != nil))
			fmt.Println(tt.wantErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitFormattingExample() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InitFormattingExample() = %v, want %v", got, tt.want)
			}
		})
	}
}

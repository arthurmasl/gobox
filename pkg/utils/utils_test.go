package utils

import "testing"

func TestSummy(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Summy(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Summy() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Summy() = %v, want %v", got, tt.want)
			}
		})
	}
}

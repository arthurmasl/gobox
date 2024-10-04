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
		{
			name:    "sum",
			args:    args{5, 1},
			want:    6,
			wantErr: false,
		},
		{
			name:    "sum2",
			args:    args{5, 15},
			want:    20,
			wantErr: false,
		},
		{
			name:    "err",
			args:    args{505, 15},
			want:    -1,
			wantErr: true,
		},
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

package crypt

import "testing"

func TestPasswordHash(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "01",
			args:    args{password: "10031234"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := PasswordHash(tt.args.password)
				if (err != nil) != tt.wantErr {
					t.Errorf("PasswordHash() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got == "" {
					t.Errorf("PasswordHash() got empty")
				}
			},
		)
	}
}

func TestPasswordVerify(t *testing.T) {
	type args struct {
		hash     string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "false",
			args: args{
				hash:     "123",
				password: "123456",
			},
			want: false,
		},
		{
			name: "true",
			args: args{
				hash:     "$2y$10$rbhshxsqufUXdViIm.P2Oerwuyo98DxSXlhKBVeVAMZa9er9v1R7i",
				password: "10031234pFs6ioSVKL",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := PasswordVerify(tt.args.hash, tt.args.password); got != tt.want {
					t.Errorf("PasswordVerify() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

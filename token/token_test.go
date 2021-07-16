package token

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

// 测试生成Token
func TestGenerateToken(t *testing.T) {
	tests := []struct {
		name    string
		claims  UserTokenClaims
		want    string
		wantErr bool
	}{
		{
			name: "01",
			claims: UserTokenClaims{
				UserID:         1,
				Mobile:         "18702982564",
				FamilyID:       1,
				StandardClaims: jwt.StandardClaims{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := GenerateToken(tt.claims)
				if (err != nil) != tt.wantErr {
					t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				fmt.Println(got)
				if got == "" {
					t.Errorf("GenerateToken() got empty")
					return
				}
			},
		)
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    *UserTokenClaims
		wantErr bool
	}{
		{
			name: "success",
			args: args{s: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJtb2JpbGUiOiIxODcwMjk4MjU2NCIsImZhbWlseV9pZCI6MX0.JBxAdmLcmZZ9yBI_7ghrd7N-vqZF_XR2yJ-0i0YRHcY"},
			want: &UserTokenClaims{
				UserID:         1,
				Mobile:         "18702982564",
				FamilyID:       1,
				StandardClaims: jwt.StandardClaims{},
			},
			wantErr: false,
		},
		{
			name:    "failed",
			args:    args{s: "111"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := ParseToken(tt.args.s)
				if (err != nil) != tt.wantErr {
					t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
					t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

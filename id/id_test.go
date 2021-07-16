package id

import (
	"testing"
)

// 测试生成ID
func TestNewID(t *testing.T) {
	tests := []struct {
		name    string
		want    int64
		wantErr bool
	}{
		{
			name:    "normal",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				_, err := NewID()
				if (err != nil) != tt.wantErr {
					t.Errorf("NewID() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
			},
		)
	}
}

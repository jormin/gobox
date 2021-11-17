package id

import (
	"testing"
)

// 测试生成ID
func TestNewID(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{
			name: "normal",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if NewID() == "" {
					t.Error("NewID() error, got empty")
					return
				}
			},
		)
	}
}

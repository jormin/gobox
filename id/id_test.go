package id

import (
	"testing"

	"github.com/jormin/golog/log"
)

// 测试生成ID
func TestNewID(t *testing.T) {
	cfg := Config{
		Host: "111.229.255.133",
		Port: 31307,
	}
	err := Init(&cfg)
	if err != nil {
		t.Fatalf("init error: %v", err)
	}
	id, err := NewID()
	if err != nil {
		t.Errorf("get id error: %v", err)
	}
	log.Info("get id success: %d", id)
}

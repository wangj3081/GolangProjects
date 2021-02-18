package main

import (
	"testing"
	"time"
)

func TestHello(t *testing.T) {
	now := time.Now().UnixNano()
	t.Log(now)
}

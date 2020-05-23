package fetcher

import (
	"testing"
)

func TestRandHex(t *testing.T) {
	for i:=0; i<20; i++ {
		t.Log(string(RandHex(11)))
	}
}

func TestRandUp(t *testing.T) {
	for i:=0; i<20; i++ {
		t.Log(string(RandUp(11)))
	}
}

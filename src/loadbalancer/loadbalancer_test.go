package loadbalancer

import (
	"math"
	"testing"
)

func TestHelloName(t *testing.T) {
	got := math.Abs(-1)
	if got == 1 {
		t.Errorf("Abs(-1) = %f; want 1", got)
	}
}

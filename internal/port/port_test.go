package port

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	min, max := 10000, 20000
	port := Generate(min, max)
	if port < min || port > max {
		t.Errorf("Generated port %d is not within the range %d-%d", port, min, max)
	}
}

func TestIsInUse(t *testing.T) {
	// This test is a bit tricky as it depends on system state
	// We'll test for a very high port that's unlikely to be in use
	port := 65000
	if isInUse(port) {
		t.Errorf("Port %d is reported as in use, but it should be free", port)
	}
}

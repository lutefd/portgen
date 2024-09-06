package app

import (
	"testing"
)

func TestGeneratePort(t *testing.T) {
	min, max := 10000, 20000
	port := GeneratePort(min, max)
	if port < min || port > max {
		t.Errorf("Generated port %d is not within the range %d-%d", port, min, max)
	}
}

func TestIsPortInUse(t *testing.T) {
	// This test is a bit tricky as it depends on system state
	// We'll test for a very high port that's unlikely to be in use
	port := 65000
	if isPortInUse(port) {
		t.Errorf("Port %d is reported as in use, but it should be free", port)
	}
}

// Mock the clipboard for testing
type mockClipboard struct {
	content string
}

func (m *mockClipboard) WriteAll(text string) error {
	m.content = text
	return nil
}

func TestCopyToClipboard(t *testing.T) {
	// Replace the actual clipboard with our mock
	oldClipboard := defaultClipboard
	mockClip := &mockClipboard{}
	defaultClipboard = mockClip
	defer func() { defaultClipboard = oldClipboard }()

	port := 12345
	err := CopyToClipboard(port)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if mockClip.content != "12345" {
		t.Errorf("Expected '12345' to be copied to clipboard, but got '%s'", mockClip.content)
	}
}

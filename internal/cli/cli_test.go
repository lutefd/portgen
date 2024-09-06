package cli

import (
	"bytes"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{"portgen", "--test"}
	Execute()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	port, err := strconv.Atoi(strings.TrimSpace(output))
	if err != nil {
		t.Fatalf("Expected a port number, got: %s", output)
	}

	if port < 10000 || port > 65535 {
		t.Errorf("Port %d is out of the expected range (10000-65535)", port)
	}
}

func TestExecuteWithFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		minPort  int
		maxPort  int
		wantCopy bool
	}{
		{"Default", []string{"--test"}, 10000, 65535, false},
		{"Custom Range", []string{"--test", "--min", "20000", "--max", "30000"}, 20000, 30000, false},
		{"With Copy", []string{"--test", "--copy"}, 10000, 65535, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			os.Args = append([]string{"portgen"}, tt.args...)
			Execute()

			w.Close()
			os.Stdout = oldStdout

			var buf bytes.Buffer
			buf.ReadFrom(r)
			output := strings.TrimSpace(buf.String())

			lines := strings.Split(output, "\n")
			for i, line := range lines {
				lines[i] = strings.TrimSpace(line)
			}

			var nonEmptyLines []string
			for _, line := range lines {
				if line != "" {
					nonEmptyLines = append(nonEmptyLines, line)
				}
			}

			if len(nonEmptyLines) == 0 {
				t.Fatal("Expected at least one line of output, got none")
			}

			portStr := nonEmptyLines[len(nonEmptyLines)-1]

			port, err := strconv.Atoi(portStr)
			if err != nil {
				t.Fatalf("Expected a port number, got: %s", portStr)
			}

			if port < tt.minPort || port > tt.maxPort {
				t.Errorf("Port %d is out of the expected range (%d-%d)", port, tt.minPort, tt.maxPort)
			}
		})
	}
}

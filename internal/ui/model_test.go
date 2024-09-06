package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestInitialModel(t *testing.T) {
	m := InitialModel(10000, 20000, true)
	if m.minPort != 10000 || m.maxPort != 20000 || !m.copyToClipboard || m.state != stateNormal {
		t.Error("InitialModel did not set fields correctly")
	}
}

func TestModelUpdate(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		initialCopy    bool
		wantPort       bool
		wantState      modelState
		wantCopy       bool
		wantCopyChange bool
	}{
		{"Generate (empty input)", "", true, true, stateNormal, true, false},
		{"Generate command", "generate", true, true, stateNormal, true, false},
		{"Copy command", "copy", false, false, stateNormal, true, true},
		{"Help command", "help", false, false, stateHelp, false, false},
		{"Unknown command", "unknown", false, false, stateNormal, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := InitialModel(10000, 20000, tt.initialCopy)
			m.port = 12345

			m.textInput.SetValue(tt.input)
			newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
			updatedModel := newModel.(model)

			if tt.wantPort && updatedModel.port == 0 {
				t.Error("Model did not generate a port when expected")
			}

			if !tt.wantPort && tt.input == "" && updatedModel.port != 0 {
				t.Error("Model generated a port when not expected")
			}

			if updatedModel.state != tt.wantState {
				t.Errorf("Model state is %v, want %v", updatedModel.state, tt.wantState)
			}

			if updatedModel.copyToClipboard != tt.wantCopy {
				t.Errorf("Copy to clipboard is %v, want %v", updatedModel.copyToClipboard, tt.wantCopy)
			}

			if tt.wantCopyChange && updatedModel.copyToClipboard == m.copyToClipboard {
				t.Error("Copy to clipboard flag did not change when expected")
			}

			if tt.input == "unknown" && updatedModel.err == nil {
				t.Error("Model did not set error for unknown command")
			}

			if tt.wantState == stateHelp {
				newModel, _ = updatedModel.Update(tea.KeyMsg{Type: tea.KeyEnter})
				finalModel := newModel.(model)
				if finalModel.state != stateNormal {
					t.Error("Model did not return to normal state after displaying help")
				}
			}
		})
	}
}

func TestModelView(t *testing.T) {
	m := InitialModel(10000, 20000, true)
	view := m.View()
	if !strings.Contains(view, "Portgen") {
		t.Error("Model View did not contain expected title")
	}

	m.state = stateHelp
	helpView := m.View()
	if !strings.Contains(helpView, "Help") || !strings.Contains(helpView, "Commands:") {
		t.Error("Help view did not contain expected content")
	}

	m.err = &errorString{"Test error"}
	errorView := m.View()
	if !strings.Contains(errorView, "Error: Test error") {
		t.Error("Error view did not contain expected error message")
	}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

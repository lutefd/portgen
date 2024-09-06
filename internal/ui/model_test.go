package ui

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestInitialModel(t *testing.T) {
	m := InitialModel(10000, 20000, true)
	if m.minPort != 10000 || m.maxPort != 20000 || !m.copyToClipboard {
		t.Error("InitialModel did not set fields correctly")
	}
}

func TestModelUpdate(t *testing.T) {
	m := InitialModel(10000, 20000, true)
	newModel, _ := m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	updatedModel := newModel.(model)
	if updatedModel.port == 0 {
		t.Error("Model did not generate a port on Enter key")
	}
}

func TestModelView(t *testing.T) {
	m := InitialModel(10000, 20000, true)
	view := m.View()
	if view == "" {
		t.Error("Model View returned an empty string")
	}
}

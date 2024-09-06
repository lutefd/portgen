package ui

import (
	"fmt"
	"strconv"

	"github.com/Lutefd/portgen/internal/port"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	minPort         int
	maxPort         int
	copyToClipboard bool
	port            int
	textInput       textinput.Model
	err             error
}

func InitialModel(minPort, maxPort int, copyToClipboard bool) model {
	ti := textinput.New()
	ti.Placeholder = "Press enter to generate a port"
	ti.Focus()

	return model{
		minPort:         minPort,
		maxPort:         maxPort,
		copyToClipboard: copyToClipboard,
		textInput:       ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.port = port.Generate(m.minPort, m.maxPort)
			if m.copyToClipboard {
				clipboard.WriteAll(strconv.Itoa(m.port))
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v", m.err)
	}

	var s string
	if m.port != 0 {
		s = PortStyle.Render(fmt.Sprintf("Generated Port: %d", m.port))
		if m.copyToClipboard {
			s += "\n" + InfoStyle.Render("(Copied to clipboard)")
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		TitleStyle.Render("Random Port Generator"),
		InputStyle.Render(m.textInput.View()),
		s,
		InfoStyle.Render("(esc to quit)"),
	)
}

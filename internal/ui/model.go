package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Lutefd/portgen/internal/port"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type modelState int

const (
	stateNormal modelState = iota
	stateHelp
)

type model struct {
	minPort         int
	maxPort         int
	copyToClipboard bool
	port            int
	textInput       textinput.Model
	err             error
	state           modelState
}

func InitialModel(minPort, maxPort int, copyToClipboard bool) model {
	ti := textinput.New()
	ti.Placeholder = "Enter command (generate, copy, help) or press enter"
	ti.Focus()

	return model{
		minPort:         minPort,
		maxPort:         maxPort,
		copyToClipboard: copyToClipboard,
		textInput:       ti,
		state:           stateNormal,
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
			if m.state == stateHelp {
				m.state = stateNormal
				return m, nil
			}
			command := strings.TrimSpace(strings.ToLower(m.textInput.Value()))
			switch command {
			case "", "generate":
				m.port = port.Generate(m.minPort, m.maxPort)
				if m.copyToClipboard {
					clipboard.WriteAll(strconv.Itoa(m.port))
				}
			case "copy":
				if m.port != 0 {
					clipboard.WriteAll(strconv.Itoa(m.port))
					m.copyToClipboard = true
				}
			case "help":
				m.state = stateHelp
			default:
				m.err = fmt.Errorf("unknown command: %s", command)
			}
			m.textInput.SetValue("")
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
		return fmt.Sprintf("Error: %v\n\nPress Enter to continue", m.err)
	}

	if m.state == stateHelp {
		return lipgloss.JoinVertical(lipgloss.Left,
			TitleStyle.Render("Help"),
			HelpDescStyle.Render("Commands:"),
			HelpCommandStyle.Render("  generate: ")+HelpDescStyle.Render("Generate a new port"),
			HelpCommandStyle.Render("  copy:     ")+HelpDescStyle.Render("Copy the current port to clipboard"),
			HelpCommandStyle.Render("  help:     ")+HelpDescStyle.Render("Show this help message"),
			"",
			InfoStyle.Render("Press Enter to return to the main screen"),
			InfoStyle.Render("(esc to quit)"),
		)
	}

	var s string
	if m.port != 0 {
		s = PortStyle.Render(fmt.Sprintf("Generated Port: %d", m.port))
		if m.copyToClipboard {
			s += "\n" + InfoStyle.Render("(Copied to clipboard)")
		}
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		TitleStyle.Render("Portgen"),
		InputStyle.Render(m.textInput.View()),
		s,
		InfoStyle.Render("(type 'help' for commands, esc to quit)"),
	)
}

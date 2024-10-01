package ui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lutefd/portgen/internal/port"
)

type modelState int

const (
	stateNormal modelState = iota
	stateHelp
	stateError
)

type model struct {
	minPort         int
	maxPort         int
	copyToClipboard bool
	port            int
	textInput       textinput.Model
	err             error
	state           modelState
	tempMessage     string
	tempMessageTime time.Time
}

func InitialModel(minPort, maxPort int, copyToClipboard bool) model {
	ti := textinput.New()
	ti.Placeholder = "Enter command (generate, copy, toggle, help) or press enter"
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
			if m.state == stateHelp || m.state == stateError {
				m.state = stateNormal
				m.err = nil
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
					m.tempMessage = fmt.Sprintf("Port %d copied to clipboard", m.port)
					m.tempMessageTime = time.Now()
					m.textInput.SetValue("")
					return m, clearTempMessageAfter(2 * time.Second)
				}
			case "c":
				if m.port != 0 {
					clipboard.WriteAll(strconv.Itoa(m.port))
					m.tempMessage = fmt.Sprintf("Port %d copied to clipboard", m.port)
					m.tempMessageTime = time.Now()
					m.textInput.SetValue("")
					return m, clearTempMessageAfter(2 * time.Second)
				}
			case "toggle":
				m.copyToClipboard = !m.copyToClipboard
			case "t":
				m.copyToClipboard = !m.copyToClipboard
			case "help":
				m.state = stateHelp
			default:
				m.state = stateError
				m.err = fmt.Errorf("unknown command: %s", command)
			}
			m.textInput.SetValue("")
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case error:
		m.state = stateError
		m.err = msg
		return m, nil
	case clearTempMessageMsg:
		m.tempMessage = ""
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.state == stateHelp {
		return lipgloss.JoinVertical(lipgloss.Left,
			TitleStyle.Render("Help"),
			HelpDescStyle.Render("Commands:"),
			HelpCommandStyle.Render("  generate: ")+HelpDescStyle.Render("Generate a new port"),
			HelpCommandStyle.Render("  copy:     ")+HelpDescStyle.Render("Copy the current port to clipboard, you can also use 'c'"),
			HelpCommandStyle.Render("  toggle:   ")+HelpDescStyle.Render("Toggle clipboard mode, you can also use 't'"),
			HelpCommandStyle.Render("  help:     ")+HelpDescStyle.Render("Show this help message"),
			"",
			InfoStyle.Render("Press Enter to return to the main screen"),
			InfoStyle.Render("(esc to quit)"),
		)
	}
	if m.state == stateError {
		return lipgloss.JoinVertical(lipgloss.Left,
			TitleStyle.Render("Error"),
			ErrorDescStyle.Render(m.err.Error()),
			InfoStyle.Render("Press Enter to return to the main screen"),
			InfoStyle.Render("(esc to quit)"),
		)
	}

	clipboardStatus := fmt.Sprintf(" [Clipboard: %s]", onOffText(m.copyToClipboard))
	clipboardStatusStyled := ClipboardStatusStyle.Render(clipboardStatus)
	if m.copyToClipboard {
		clipboardStatusStyled = ClipboardOnStyle.Render(clipboardStatus)
	} else {
		clipboardStatusStyled = ClipboardOffStyle.Render(clipboardStatus)
	}

	title := lipgloss.JoinHorizontal(lipgloss.Center,
		TitleStyle.Render("Portgen"),
		lipgloss.NewStyle().MarginLeft(1).Render("—"),
		clipboardStatusStyled,
	)

	var s string
	if m.port != 0 {
		s = PortStyle.Render(fmt.Sprintf("Generated Port: %d", m.port))
	}

	if m.tempMessage != "" {
		s += "\n" + TempMessageStyle.Render(m.tempMessage)
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		title,
		InputStyle.Render(m.textInput.View()),
		s,
		InfoStyle.Render("(type 'help' for commands, esc to quit)"),
	)
}

func onOffText(b bool) string {
	if b {
		return "ON"
	}
	return "OFF"
}

type clearTempMessageMsg struct{}

func clearTempMessageAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return clearTempMessageMsg{}
	})
}

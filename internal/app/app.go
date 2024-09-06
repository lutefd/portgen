package app

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/Lutefd/portgen/internal/ui"
	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// clipboardWriter is an interface that allows us to mock the clipboard
type clipboardWriter interface {
	WriteAll(text string) error
}

// realClipboard wraps the actual clipboard package
type realClipboard struct{}

func (rc realClipboard) WriteAll(text string) error {
	return clipboard.WriteAll(text)
}

// defaultClipboard is the clipboard implementation we'll use by default
var defaultClipboard clipboardWriter = realClipboard{}

func GeneratePort(min, max int) int {
	for {
		port := rand.Intn(max-min+1) + min
		if !isPortInUse(port) {
			return port
		}
	}
}

func isPortInUse(port int) bool {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return true
	}
	listener.Close()
	return false
}

func CopyToClipboard(port int) error {
	return defaultClipboard.WriteAll(strconv.Itoa(port))
}

func RunInteractiveMode(minPort, maxPort int, copyToClipboard bool) {
	p := tea.NewProgram(ui.InitialModel(minPort, maxPort, copyToClipboard))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
	}
}

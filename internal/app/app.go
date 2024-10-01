package app

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/lutefd/portgen/internal/ui"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

type clipboardWriter interface {
	WriteAll(text string) error
}

type realClipboard struct{}

func (rc realClipboard) WriteAll(text string) error {
	return clipboard.WriteAll(text)
}

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

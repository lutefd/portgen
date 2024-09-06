package ui

import (
	"strings"
	"testing"
)

func TestGetLongDescription(t *testing.T) {
	desc := GetLongDescription()
	if !strings.Contains(desc, "Portgen") {
		t.Error("Long description does not contain expected text")
	}
}

func TestGetUsageTemplate(t *testing.T) {
	usage := GetUsageTemplate()
	if !strings.Contains(usage, "Usage:") || !strings.Contains(usage, "Flags:") {
		t.Error("Usage template does not contain expected sections")
	}
}

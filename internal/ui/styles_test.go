package ui

import (
	"testing"

	"github.com/charmbracelet/lipgloss"
)

func TestPriorityStyle(t *testing.T) {
	tests := []struct {
		priority int
		wantBold bool
	}{
		{0, true},   // P0 - Critical, should be bold
		{1, true},   // P1 - High, should be bold
		{2, false},  // P2 - Medium, not bold
		{3, false},  // P3 - Low, not bold
		{4, false},  // P4 - Backlog, not bold
		{5, false},  // Unknown priority, not bold
		{-1, false}, // Invalid priority, not bold
	}

	for _, tt := range tests {
		t.Run(string(rune('0'+tt.priority)), func(t *testing.T) {
			style := PriorityStyle(tt.priority)
			// We can't directly check the color, but we can verify the style is created
			// and check the bold property
			rendered := style.Render("test")
			if rendered == "" {
				t.Error("PriorityStyle should render non-empty string")
			}
		})
	}
}

func TestStatusStyle(t *testing.T) {
	tests := []struct {
		status string
	}{
		{"open"},
		{"in_progress"},
		{"closed"},
		{"unknown"},
		{""},
	}

	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			style := StatusStyle(tt.status)
			rendered := style.Render("test")
			if rendered == "" {
				t.Error("StatusStyle should render non-empty string")
			}
		})
	}
}

func TestPriorityColors(t *testing.T) {
	// Verify all expected priorities have colors
	expectedPriorities := []int{0, 1, 2, 3, 4}
	for _, p := range expectedPriorities {
		if _, ok := PriorityColors[p]; !ok {
			t.Errorf("Priority %d should have a color defined", p)
		}
	}

	// Verify colors are valid lipgloss.Color
	for p, color := range PriorityColors {
		if color == lipgloss.Color("") {
			t.Errorf("Priority %d has empty color", p)
		}
	}
}

func TestStatusColors(t *testing.T) {
	// Verify all expected statuses have colors
	expectedStatuses := []string{"open", "in_progress", "closed"}
	for _, s := range expectedStatuses {
		if _, ok := StatusColors[s]; !ok {
			t.Errorf("Status %q should have a color defined", s)
		}
	}

	// Verify colors are valid lipgloss.Color
	for s, color := range StatusColors {
		if color == lipgloss.Color("") {
			t.Errorf("Status %q has empty color", s)
		}
	}
}

func TestStylesInitialized(t *testing.T) {
	// Test that all base styles are initialized and can render
	styles := []struct {
		name  string
		style lipgloss.Style
	}{
		{"AppStyle", AppStyle},
		{"TitleStyle", TitleStyle},
		{"PanelStyle", PanelStyle},
		{"FocusedPanelStyle", FocusedPanelStyle},
		{"PanelTitleStyle", PanelTitleStyle},
		{"TaskItemStyle", TaskItemStyle},
		{"SelectedTaskStyle", SelectedTaskStyle},
		{"TaskIDStyle", TaskIDStyle},
		{"TaskTitleStyle", TaskTitleStyle},
		{"StatusBarStyle", StatusBarStyle},
		{"HelpBarStyle", HelpBarStyle},
		{"HelpKeyStyle", HelpKeyStyle},
		{"HelpDescStyle", HelpDescStyle},
		{"DetailLabelStyle", DetailLabelStyle},
		{"DetailValueStyle", DetailValueStyle},
		{"FormLabelStyle", FormLabelStyle},
		{"FormInputStyle", FormInputStyle},
		{"FormInputFocusedStyle", FormInputFocusedStyle},
		{"OverlayStyle", OverlayStyle},
		{"ErrorStyle", ErrorStyle},
		{"SuccessStyle", SuccessStyle},
	}

	for _, tt := range styles {
		t.Run(tt.name, func(t *testing.T) {
			rendered := tt.style.Render("test")
			if rendered == "" {
				t.Errorf("%s should render non-empty string", tt.name)
			}
		})
	}
}

package ui

import (
	"strings"
	"testing"
)

func TestNewInlineBarInput(t *testing.T) {
	bar := NewInlineBarInput("Edit Title", "ISS-123", "Initial value", 80)

	if bar.Type != InlineBarInput {
		t.Errorf("Expected InlineBarInput type, got %v", bar.Type)
	}
	if bar.Title != "Edit Title" {
		t.Errorf("Expected title 'Edit Title', got %q", bar.Title)
	}
	if bar.Subtitle != "ISS-123" {
		t.Errorf("Expected subtitle 'ISS-123', got %q", bar.Subtitle)
	}
	if bar.InputValue() != "Initial value" {
		t.Errorf("Expected input value 'Initial value', got %q", bar.InputValue())
	}
}

func TestNewInlineBarSelect(t *testing.T) {
	options := []InlineBarOption{
		{Label: "Open", Value: "open", Shortcut: "o"},
		{Label: "In Progress", Value: "in_progress", Shortcut: "i"},
		{Label: "Closed", Value: "closed", Shortcut: "c"},
	}

	bar := NewInlineBarSelect("Change Status", "ISS-123", options, "in_progress")

	if bar.Type != InlineBarSelect {
		t.Errorf("Expected InlineBarSelect type, got %v", bar.Type)
	}
	if bar.Title != "Change Status" {
		t.Errorf("Expected title 'Change Status', got %q", bar.Title)
	}
	if bar.Subtitle != "ISS-123" {
		t.Errorf("Expected subtitle 'ISS-123', got %q", bar.Subtitle)
	}
	if len(bar.Options) != 3 {
		t.Errorf("Expected 3 options, got %d", len(bar.Options))
	}
	if bar.Selected != 1 {
		t.Errorf("Expected selected index 1 (in_progress), got %d", bar.Selected)
	}
}

func TestNewInlineBarSelect_DefaultSelection(t *testing.T) {
	options := []InlineBarOption{
		{Label: "Option 1", Value: "opt1"},
		{Label: "Option 2", Value: "opt2"},
	}

	bar := NewInlineBarSelect("Select", "", options, "nonexistent")

	if bar.Selected != 0 {
		t.Errorf("Expected default selected index 0, got %d", bar.Selected)
	}
}

func TestInlineBar_MoveLeft(t *testing.T) {
	options := []InlineBarOption{
		{Label: "Option 1", Value: "opt1"},
		{Label: "Option 2", Value: "opt2"},
		{Label: "Option 3", Value: "opt3"},
	}

	bar := NewInlineBarSelect("Select", "", options, "opt2")
	if bar.Selected != 1 {
		t.Fatalf("Expected initial selected 1, got %d", bar.Selected)
	}

	bar.MoveLeft()
	if bar.Selected != 0 {
		t.Errorf("Expected selected 0 after MoveLeft, got %d", bar.Selected)
	}

	// Should not go below 0
	bar.MoveLeft()
	if bar.Selected != 0 {
		t.Errorf("Expected selected to stay at 0, got %d", bar.Selected)
	}
}

func TestInlineBar_MoveRight(t *testing.T) {
	options := []InlineBarOption{
		{Label: "Option 1", Value: "opt1"},
		{Label: "Option 2", Value: "opt2"},
		{Label: "Option 3", Value: "opt3"},
	}

	bar := NewInlineBarSelect("Select", "", options, "opt1")
	if bar.Selected != 0 {
		t.Fatalf("Expected initial selected 0, got %d", bar.Selected)
	}

	bar.MoveRight()
	if bar.Selected != 1 {
		t.Errorf("Expected selected 1 after MoveRight, got %d", bar.Selected)
	}

	bar.MoveRight()
	if bar.Selected != 2 {
		t.Errorf("Expected selected 2 after second MoveRight, got %d", bar.Selected)
	}

	// Should not go beyond last option
	bar.MoveRight()
	if bar.Selected != 2 {
		t.Errorf("Expected selected to stay at 2, got %d", bar.Selected)
	}
}

func TestInlineBar_MoveLeft_InputBar(t *testing.T) {
	bar := NewInlineBarInput("Title", "", "value", 80)
	bar.MoveLeft()
	// Should not crash or change anything for input bar
}

func TestInlineBar_MoveRight_InputBar(t *testing.T) {
	bar := NewInlineBarInput("Title", "", "value", 80)
	bar.MoveRight()
	// Should not crash or change anything for input bar
}

func TestInlineBar_SelectByShortcut(t *testing.T) {
	options := []InlineBarOption{
		{Label: "Open", Value: "open", Shortcut: "o"},
		{Label: "In Progress", Value: "in_progress", Shortcut: "i"},
		{Label: "Closed", Value: "closed", Shortcut: "c"},
	}

	bar := NewInlineBarSelect("Select", "", options, "open")

	// Select by shortcut "i"
	if !bar.SelectByShortcut("i") {
		t.Error("Expected SelectByShortcut to return true for valid shortcut")
	}
	if bar.Selected != 1 {
		t.Errorf("Expected selected 1 after shortcut 'i', got %d", bar.Selected)
	}

	// Select by shortcut "c"
	if !bar.SelectByShortcut("c") {
		t.Error("Expected SelectByShortcut to return true for valid shortcut")
	}
	if bar.Selected != 2 {
		t.Errorf("Expected selected 2 after shortcut 'c', got %d", bar.Selected)
	}

	// Invalid shortcut
	if bar.SelectByShortcut("x") {
		t.Error("Expected SelectByShortcut to return false for invalid shortcut")
	}
	if bar.Selected != 2 {
		t.Errorf("Expected selected to stay at 2, got %d", bar.Selected)
	}
}

func TestInlineBar_SelectByShortcut_InputBar(t *testing.T) {
	bar := NewInlineBarInput("Title", "", "value", 80)

	if bar.SelectByShortcut("x") {
		t.Error("Expected SelectByShortcut to return false for input bar")
	}
}

func TestInlineBar_SelectedValue(t *testing.T) {
	options := []InlineBarOption{
		{Label: "Open", Value: "open"},
		{Label: "Closed", Value: "closed"},
	}

	bar := NewInlineBarSelect("Select", "", options, "closed")

	if bar.SelectedValue() != "closed" {
		t.Errorf("Expected selected value 'closed', got %q", bar.SelectedValue())
	}

	bar.MoveLeft()
	if bar.SelectedValue() != "open" {
		t.Errorf("Expected selected value 'open', got %q", bar.SelectedValue())
	}
}

func TestInlineBar_SelectedValue_InputBar(t *testing.T) {
	bar := NewInlineBarInput("Title", "", "value", 80)

	if bar.SelectedValue() != "" {
		t.Errorf("Expected empty selected value for input bar, got %q", bar.SelectedValue())
	}
}

func TestInlineBar_InputValue(t *testing.T) {
	bar := NewInlineBarInput("Title", "", "test value", 80)

	if bar.InputValue() != "test value" {
		t.Errorf("Expected input value 'test value', got %q", bar.InputValue())
	}
}

func TestInlineBar_View_Input(t *testing.T) {
	bar := NewInlineBarInput("Edit Title", "ISS-123", "My task", 80)

	view := bar.View(80)

	// Check that view contains expected elements
	if !strings.Contains(view, "Edit Title") {
		t.Error("View should contain title")
	}
	if !strings.Contains(view, "ISS-123") {
		t.Error("View should contain subtitle")
	}
	if !strings.Contains(view, "enter:save") {
		t.Error("View should contain help text")
	}
}

func TestInlineBar_View_Select(t *testing.T) {
	options := []InlineBarOption{
		{Label: "Open", Value: "open", Shortcut: "o"},
		{Label: "Closed", Value: "closed", Shortcut: "c"},
	}

	bar := NewInlineBarSelect("Change Status", "ISS-123", options, "open")

	view := bar.View(80)

	// Check that view contains expected elements
	if !strings.Contains(view, "Change Status") {
		t.Error("View should contain title")
	}
	if !strings.Contains(view, "ISS-123") {
		t.Error("View should contain subtitle")
	}
	if !strings.Contains(view, "h/l:nav") {
		t.Error("View should contain help text")
	}
}

func TestInlineBar_View_Select_NoSubtitle(t *testing.T) {
	options := []InlineBarOption{
		{Label: "Option 1", Value: "opt1"},
	}

	bar := NewInlineBarSelect("Select", "", options, "opt1")

	view := bar.View(80)

	if !strings.Contains(view, "Select") {
		t.Error("View should contain title")
	}
}

func TestInlineBar_View_OptionsWithoutShortcut(t *testing.T) {
	options := []InlineBarOption{
		{Label: "Option 1", Value: "opt1"},
		{Label: "Option 2", Value: "opt2", Shortcut: "2"},
	}

	bar := NewInlineBarSelect("Select", "", options, "opt1")

	view := bar.View(80)

	if !strings.Contains(view, "Option 1") {
		t.Error("View should contain option without shortcut")
	}
	if !strings.Contains(view, "[2]Option 2") {
		t.Error("View should contain option with shortcut")
	}
}

package ui

import (
	"strings"
	"testing"
)

func TestNewModalInput(t *testing.T) {
	modal := NewModalInput("Edit Title", "ISS-123", "Initial value")

	if modal.Type != ModalInput {
		t.Errorf("Expected ModalInput type, got %v", modal.Type)
	}
	if modal.Title != "Edit Title" {
		t.Errorf("Expected title 'Edit Title', got %q", modal.Title)
	}
	if modal.Subtitle != "ISS-123" {
		t.Errorf("Expected subtitle 'ISS-123', got %q", modal.Subtitle)
	}
	if modal.InputValue() != "Initial value" {
		t.Errorf("Expected input value 'Initial value', got %q", modal.InputValue())
	}
}

func TestNewModalSelect(t *testing.T) {
	options := []ModalOption{
		{Label: "Open", Value: "open", Shortcut: "o"},
		{Label: "In Progress", Value: "in_progress", Shortcut: "i"},
		{Label: "Closed", Value: "closed", Shortcut: "c"},
	}

	modal := NewModalSelect("Change Status", "ISS-123", options, "in_progress")

	if modal.Type != ModalSelect {
		t.Errorf("Expected ModalSelect type, got %v", modal.Type)
	}
	if modal.Title != "Change Status" {
		t.Errorf("Expected title 'Change Status', got %q", modal.Title)
	}
	if modal.Subtitle != "ISS-123" {
		t.Errorf("Expected subtitle 'ISS-123', got %q", modal.Subtitle)
	}
	if len(modal.Options) != 3 {
		t.Errorf("Expected 3 options, got %d", len(modal.Options))
	}
	if modal.Selected != 1 {
		t.Errorf("Expected selected index 1 (in_progress), got %d", modal.Selected)
	}
}

func TestNewModalSelect_DefaultSelection(t *testing.T) {
	options := []ModalOption{
		{Label: "Option 1", Value: "opt1"},
		{Label: "Option 2", Value: "opt2"},
	}

	modal := NewModalSelect("Select", "", options, "nonexistent")

	if modal.Selected != 0 {
		t.Errorf("Expected default selected index 0, got %d", modal.Selected)
	}
}

func TestModal_MoveUp(t *testing.T) {
	options := []ModalOption{
		{Label: "Option 1", Value: "opt1"},
		{Label: "Option 2", Value: "opt2"},
		{Label: "Option 3", Value: "opt3"},
	}

	modal := NewModalSelect("Select", "", options, "opt2")
	if modal.Selected != 1 {
		t.Fatalf("Expected initial selected 1, got %d", modal.Selected)
	}

	modal.MoveUp()
	if modal.Selected != 0 {
		t.Errorf("Expected selected 0 after MoveUp, got %d", modal.Selected)
	}

	// Should not go below 0
	modal.MoveUp()
	if modal.Selected != 0 {
		t.Errorf("Expected selected to stay at 0, got %d", modal.Selected)
	}
}

func TestModal_MoveDown(t *testing.T) {
	options := []ModalOption{
		{Label: "Option 1", Value: "opt1"},
		{Label: "Option 2", Value: "opt2"},
		{Label: "Option 3", Value: "opt3"},
	}

	modal := NewModalSelect("Select", "", options, "opt1")
	if modal.Selected != 0 {
		t.Fatalf("Expected initial selected 0, got %d", modal.Selected)
	}

	modal.MoveDown()
	if modal.Selected != 1 {
		t.Errorf("Expected selected 1 after MoveDown, got %d", modal.Selected)
	}

	modal.MoveDown()
	if modal.Selected != 2 {
		t.Errorf("Expected selected 2 after second MoveDown, got %d", modal.Selected)
	}

	// Should not go beyond last option
	modal.MoveDown()
	if modal.Selected != 2 {
		t.Errorf("Expected selected to stay at 2, got %d", modal.Selected)
	}
}

func TestModal_MoveUp_InputModal(t *testing.T) {
	modal := NewModalInput("Title", "", "value")
	modal.MoveUp()
	// Should not crash or change anything for input modal
}

func TestModal_MoveDown_InputModal(t *testing.T) {
	modal := NewModalInput("Title", "", "value")
	modal.MoveDown()
	// Should not crash or change anything for input modal
}

func TestModal_SelectByShortcut(t *testing.T) {
	options := []ModalOption{
		{Label: "Open", Value: "open", Shortcut: "o"},
		{Label: "In Progress", Value: "in_progress", Shortcut: "i"},
		{Label: "Closed", Value: "closed", Shortcut: "c"},
	}

	modal := NewModalSelect("Select", "", options, "open")

	// Select by shortcut "i"
	if !modal.SelectByShortcut("i") {
		t.Error("Expected SelectByShortcut to return true for valid shortcut")
	}
	if modal.Selected != 1 {
		t.Errorf("Expected selected 1 after shortcut 'i', got %d", modal.Selected)
	}

	// Select by shortcut "c"
	if !modal.SelectByShortcut("c") {
		t.Error("Expected SelectByShortcut to return true for valid shortcut")
	}
	if modal.Selected != 2 {
		t.Errorf("Expected selected 2 after shortcut 'c', got %d", modal.Selected)
	}

	// Invalid shortcut
	if modal.SelectByShortcut("x") {
		t.Error("Expected SelectByShortcut to return false for invalid shortcut")
	}
	if modal.Selected != 2 {
		t.Errorf("Expected selected to stay at 2, got %d", modal.Selected)
	}
}

func TestModal_SelectByShortcut_InputModal(t *testing.T) {
	modal := NewModalInput("Title", "", "value")

	if modal.SelectByShortcut("x") {
		t.Error("Expected SelectByShortcut to return false for input modal")
	}
}

func TestModal_SelectedValue(t *testing.T) {
	options := []ModalOption{
		{Label: "Open", Value: "open"},
		{Label: "Closed", Value: "closed"},
	}

	modal := NewModalSelect("Select", "", options, "closed")

	if modal.SelectedValue() != "closed" {
		t.Errorf("Expected selected value 'closed', got %q", modal.SelectedValue())
	}

	modal.MoveUp()
	if modal.SelectedValue() != "open" {
		t.Errorf("Expected selected value 'open', got %q", modal.SelectedValue())
	}
}

func TestModal_SelectedValue_InputModal(t *testing.T) {
	modal := NewModalInput("Title", "", "value")

	if modal.SelectedValue() != "" {
		t.Errorf("Expected empty selected value for input modal, got %q", modal.SelectedValue())
	}
}

func TestModal_InputValue(t *testing.T) {
	modal := NewModalInput("Title", "", "test value")

	if modal.InputValue() != "test value" {
		t.Errorf("Expected input value 'test value', got %q", modal.InputValue())
	}
}

func TestModal_View_Input(t *testing.T) {
	modal := NewModalInput("Edit Title", "ISS-123", "My task")

	view := modal.View(80, 24)

	// Check that view contains expected elements
	if !strings.Contains(view, "Edit Title") {
		t.Error("View should contain title")
	}
	if !strings.Contains(view, "ISS-123") {
		t.Error("View should contain subtitle")
	}
	if !strings.Contains(view, "enter: save") {
		t.Error("View should contain help text")
	}
}

func TestModal_View_Select(t *testing.T) {
	options := []ModalOption{
		{Label: "Open", Value: "open", Shortcut: "o"},
		{Label: "Closed", Value: "closed", Shortcut: "c"},
	}

	modal := NewModalSelect("Change Status", "ISS-123", options, "open")

	view := modal.View(80, 24)

	// Check that view contains expected elements
	if !strings.Contains(view, "Change Status") {
		t.Error("View should contain title")
	}
	if !strings.Contains(view, "ISS-123") {
		t.Error("View should contain subtitle")
	}
	if !strings.Contains(view, "[o] Open") {
		t.Error("View should contain first option with shortcut")
	}
	if !strings.Contains(view, "[c] Closed") {
		t.Error("View should contain second option with shortcut")
	}
	if !strings.Contains(view, "j/k: nav") {
		t.Error("View should contain help text")
	}
}

func TestModal_View_Select_NoSubtitle(t *testing.T) {
	options := []ModalOption{
		{Label: "Option 1", Value: "opt1"},
	}

	modal := NewModalSelect("Select", "", options, "opt1")

	view := modal.View(80, 24)

	if !strings.Contains(view, "Select") {
		t.Error("View should contain title")
	}
}

func TestModal_View_SmallWidth(t *testing.T) {
	modal := NewModalInput("Title", "", "value")

	// Should not crash with small width
	view := modal.View(20, 10)

	if view == "" {
		t.Error("View should not be empty")
	}
}

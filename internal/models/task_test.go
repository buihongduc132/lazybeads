package models

import (
	"path/filepath"
	"testing"
)

func TestTask_PriorityString(t *testing.T) {
	tests := []struct {
		priority int
		expected string
	}{
		{0, "P0"},
		{1, "P1"},
		{2, "P2"},
		{3, "P3"},
		{4, "P4"},
		{5, "P?"},
		{-1, "P?"},
		{100, "P?"},
	}

	for _, tt := range tests {
		task := Task{Priority: tt.priority}
		if got := task.PriorityString(); got != tt.expected {
			t.Errorf("PriorityString(%d) = %q, want %q", tt.priority, got, tt.expected)
		}
	}
}

func TestTask_StatusIcon(t *testing.T) {
	tests := []struct {
		status   string
		expected string
	}{
		{"open", "○"},
		{"in_progress", "◐"},
		{"closed", "●"},
		{"unknown", "?"},
		{"", "?"},
	}

	for _, tt := range tests {
		task := Task{Status: tt.status}
		if got := task.StatusIcon(); got != tt.expected {
			t.Errorf("StatusIcon(%q) = %q, want %q", tt.status, got, tt.expected)
		}
	}
}

func TestTask_IsBlocked(t *testing.T) {
	task := Task{BlockedBy: nil}
	if task.IsBlocked() {
		t.Error("Expected IsBlocked to be false for nil BlockedBy")
	}

	task = Task{BlockedBy: []string{}}
	if task.IsBlocked() {
		t.Error("Expected IsBlocked to be false for empty BlockedBy")
	}

	task = Task{BlockedBy: []string{"ISS-001"}}
	if !task.IsBlocked() {
		t.Error("Expected IsBlocked to be true for non-empty BlockedBy")
	}
}

func TestTask_FilePath(t *testing.T) {
	task := Task{ID: "ISS-123"}
	expected := filepath.Join(".beads", "issues", "ISS-123.md")
	if got := task.FilePath(); got != expected {
		t.Errorf("FilePath() = %q, want %q", got, expected)
	}
}

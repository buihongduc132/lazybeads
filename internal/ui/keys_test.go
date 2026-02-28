package ui

import (
	"testing"

	"github.com/charmbracelet/bubbles/key"
)

func TestDefaultKeyMap(t *testing.T) {
	km := DefaultKeyMap()

	// Verify key map is initialized with non-zero bindings
	// We can't easily test key.Matches without the actual key press events,
	// so we just verify the structure is created
	if km.Up.Keys() == nil {
		t.Error("Up key should be initialized")
	}
	if km.Down.Keys() == nil {
		t.Error("Down key should be initialized")
	}
	if km.Select.Keys() == nil {
		t.Error("Select key should be initialized")
	}
	if km.Add.Keys() == nil {
		t.Error("Add key should be initialized")
	}
	if km.Delete.Keys() == nil {
		t.Error("Delete key should be initialized")
	}
	if km.Quit.Keys() == nil {
		t.Error("Quit key should be initialized")
	}
	if km.Help.Keys() == nil {
		t.Error("Help key should be initialized")
	}
	if km.Filter.Keys() == nil {
		t.Error("Filter key should be initialized")
	}
}

func TestKeyMap_ShortHelp(t *testing.T) {
	km := DefaultKeyMap()
	shortHelp := km.ShortHelp()

	if len(shortHelp) != 7 {
		t.Errorf("Expected 7 short help bindings, got %d", len(shortHelp))
	}

	// Verify all bindings are initialized
	for i, binding := range shortHelp {
		if binding.Keys() == nil {
			t.Errorf("Binding at index %d is not initialized", i)
		}
	}
}

func TestKeyMap_FullHelp(t *testing.T) {
	km := DefaultKeyMap()
	fullHelp := km.FullHelp()

	// Should have 4 groups by default (no custom commands)
	if len(fullHelp) != 4 {
		t.Errorf("Expected 4 help groups, got %d", len(fullHelp))
	}

	// Check first group (navigation)
	if len(fullHelp[0]) != 6 {
		t.Errorf("Expected 6 navigation keys, got %d", len(fullHelp[0]))
	}

	// Check second group (actions)
	if len(fullHelp[1]) != 4 {
		t.Errorf("Expected 4 action keys, got %d", len(fullHelp[1]))
	}

	// Check third group (filters)
	if len(fullHelp[2]) != 4 {
		t.Errorf("Expected 4 filter keys, got %d", len(fullHelp[2]))
	}

	// Check fourth group (UI)
	if len(fullHelp[3]) != 3 {
		t.Errorf("Expected 3 UI keys, got %d", len(fullHelp[3]))
	}

	// Verify all bindings are initialized
	for i, group := range fullHelp {
		for j, binding := range group {
			if binding.Keys() == nil {
				t.Errorf("Binding at group %d, index %d is not initialized", i, j)
			}
		}
	}
}

func TestKeyMap_FullHelp_WithCustomCommands(t *testing.T) {
	km := DefaultKeyMap()
	km.CustomCommands = []key.Binding{
		key.NewBinding(key.WithKeys("1"), key.WithHelp("1", "custom 1")),
		key.NewBinding(key.WithKeys("2"), key.WithHelp("2", "custom 2")),
	}

	fullHelp := km.FullHelp()

	// Should have 5 groups now (4 default + 1 custom)
	if len(fullHelp) != 5 {
		t.Errorf("Expected 5 help groups with custom commands, got %d", len(fullHelp))
	}

	// Check last group is custom commands
	if len(fullHelp[4]) != 2 {
		t.Errorf("Expected 2 custom commands, got %d", len(fullHelp[4]))
	}
}

func TestKeyMap_FullHelp_NoCustomCommands(t *testing.T) {
	km := DefaultKeyMap()
	km.CustomCommands = nil

	fullHelp := km.FullHelp()

	// Should have 4 groups (no custom commands)
	if len(fullHelp) != 4 {
		t.Errorf("Expected 4 help groups without custom commands, got %d", len(fullHelp))
	}
}

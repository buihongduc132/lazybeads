package beads

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"lazybeads/internal/models"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Error("NewClient returned nil")
	}
}

func TestClient_IsInitialized_NotExists(t *testing.T) {
	client := NewClient()

	tmpDir, err := os.MkdirTemp("", "beads-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	oldDir, _ := os.Getwd()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp dir: %v", err)
	}
	defer os.Chdir(oldDir)

	if client.IsInitialized() {
		t.Error("Expected IsInitialized to be false in non-beads directory")
	}
}

func TestClient_IsInitialized_Exists(t *testing.T) {
	client := NewClient()

	tmpDir, err := os.MkdirTemp("", "beads-test-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	beadsDir := filepath.Join(tmpDir, ".beads")
	if err := os.Mkdir(beadsDir, 0755); err != nil {
		t.Fatalf("Failed to create .beads dir: %v", err)
	}

	oldDir, _ := os.Getwd()
	if err := os.Chdir(tmpDir); err != nil {
		t.Fatalf("Failed to change to temp dir: %v", err)
	}
	defer os.Chdir(oldDir)

	if !client.IsInitialized() {
		t.Error("Expected IsInitialized to be true in beads directory")
	}
}

func setupFakeBd(t *testing.T, responses map[string]string) (string, func()) {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "fake-bd-*")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}

	for cmd, response := range responses {
		script := filepath.Join(tmpDir, "bd-"+cmd)
		if err := os.WriteFile(script, []byte(response), 0644); err != nil {
			os.RemoveAll(tmpDir)
			t.Fatalf("Failed to write fake bd response: %v", err)
		}
	}

	mainScript := filepath.Join(tmpDir, "bd")
	mainContent := "#!/bin/sh\ncase \"$1\" in\n"
	for cmd := range responses {
		mainContent += "  " + cmd + ") cat \"" + tmpDir + "/bd-" + cmd + "\" ;;\n"
	}
	mainContent += "  *) echo '{}' ;;\nesac"
	if err := os.WriteFile(mainScript, []byte(mainContent), 0755); err != nil {
		os.RemoveAll(tmpDir)
		t.Fatalf("Failed to write fake bd script: %v", err)
	}

	oldPath := os.Getenv("PATH")
	os.Setenv("PATH", tmpDir+":"+oldPath)

	return tmpDir, func() {
		os.Setenv("PATH", oldPath)
		os.RemoveAll(tmpDir)
	}
}

func TestClient_List_Unit(t *testing.T) {
	tasks := []models.Task{
		{ID: "ISS-001", Title: "Task 1", Status: "open", Priority: 1, Type: "task"},
		{ID: "ISS-002", Title: "Task 2", Status: "closed", Priority: 2, Type: "bug"},
	}
	response, _ := json.Marshal(tasks)

	_, cleanup := setupFakeBd(t, map[string]string{"list": string(response)})
	defer cleanup()

	client := NewClient()
	result, err := client.List()
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(result))
	}
	if result[0].ID != "ISS-001" {
		t.Errorf("Expected first task ID 'ISS-001', got %q", result[0].ID)
	}
}

func TestClient_ListOpen_Unit(t *testing.T) {
	tasks := []models.Task{
		{ID: "ISS-001", Title: "Open Task", Status: "open", Priority: 1, Type: "task"},
	}
	response, _ := json.Marshal(tasks)

	_, cleanup := setupFakeBd(t, map[string]string{"list": string(response)})
	defer cleanup()

	client := NewClient()
	result, err := client.ListOpen()
	if err != nil {
		t.Fatalf("ListOpen failed: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 task, got %d", len(result))
	}
}

func TestClient_Ready_Unit(t *testing.T) {
	tasks := []models.Task{
		{ID: "ISS-001", Title: "Ready Task", Status: "open", Priority: 1, Type: "task"},
	}
	response, _ := json.Marshal(tasks)

	_, cleanup := setupFakeBd(t, map[string]string{"ready": string(response)})
	defer cleanup()

	client := NewClient()
	result, err := client.Ready()
	if err != nil {
		t.Fatalf("Ready failed: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 task, got %d", len(result))
	}
}

func TestClient_Show_Unit(t *testing.T) {
	tasks := []models.Task{
		{ID: "ISS-001", Title: "Test Task", Status: "open", Priority: 1, Type: "task"},
	}
	response, _ := json.Marshal(tasks)

	_, cleanup := setupFakeBd(t, map[string]string{"show": string(response)})
	defer cleanup()

	client := NewClient()
	result, err := client.Show("ISS-001")
	if err != nil {
		t.Fatalf("Show failed: %v", err)
	}

	if result.ID != "ISS-001" {
		t.Errorf("Expected task ID 'ISS-001', got %q", result.ID)
	}
}

func TestClient_Show_EmptyResult(t *testing.T) {
	response := "[]"

	_, cleanup := setupFakeBd(t, map[string]string{"show": response})
	defer cleanup()

	client := NewClient()
	_, err := client.Show("ISS-999")
	if err == nil {
		t.Error("Expected error for empty result")
	}
}

func TestClient_Create_Unit(t *testing.T) {
	task := models.Task{ID: "ISS-001", Title: "New Task", Status: "open", Priority: 1, Type: "task"}
	response, _ := json.Marshal(task)

	_, cleanup := setupFakeBd(t, map[string]string{"create": string(response)})
	defer cleanup()

	client := NewClient()
	result, err := client.Create(CreateOptions{
		Title:    "New Task",
		Type:     "task",
		Priority: 1,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if result.ID != "ISS-001" {
		t.Errorf("Expected task ID 'ISS-001', got %q", result.ID)
	}
}

func TestClient_Create_WithLabels(t *testing.T) {
	task := models.Task{ID: "ISS-001", Title: "New Task", Status: "open", Priority: 1, Type: "task"}
	response, _ := json.Marshal(task)

	_, cleanup := setupFakeBd(t, map[string]string{"create": string(response)})
	defer cleanup()

	client := NewClient()
	result, err := client.Create(CreateOptions{
		Title:       "New Task",
		Type:        "task",
		Priority:    1,
		Labels:      []string{"bug", "urgent"},
		Description: "Test description",
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if result.ID != "ISS-001" {
		t.Errorf("Expected task ID 'ISS-001', got %q", result.ID)
	}
}

func TestClient_Create_InvalidPriority(t *testing.T) {
	task := models.Task{ID: "ISS-001", Title: "New Task", Status: "open", Priority: 0, Type: "task"}
	response, _ := json.Marshal(task)

	_, cleanup := setupFakeBd(t, map[string]string{"create": string(response)})
	defer cleanup()

	client := NewClient()
	result, err := client.Create(CreateOptions{
		Title:    "New Task",
		Priority: -1, // Invalid, should not be included
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	if result.ID != "ISS-001" {
		t.Errorf("Expected task ID 'ISS-001', got %q", result.ID)
	}
}

func skipIfNoBeads(t *testing.T) {
	t.Helper()
	if os.Getenv("BEADS_INTEGRATION") == "" {
		t.Skip("BEADS_INTEGRATION is not set, skipping integration test")
	}
	if _, err := os.Stat(".beads"); os.IsNotExist(err) {
		// Try parent directories up to 3 levels
		for _, dir := range []string{"..", "../..", "../../.."} {
			if _, err := os.Stat(dir + "/.beads"); err == nil {
				if err := os.Chdir(dir); err == nil {
					return
				}
			}
		}
		t.Skip("No .beads directory found, skipping integration test")
	}
}

func TestClient_IsInitialized(t *testing.T) {
	skipIfNoBeads(t)
	client := NewClient()

	if !client.IsInitialized() {
		t.Error("Expected IsInitialized to return true in beads directory")
	}
}

func TestClient_List(t *testing.T) {
	skipIfNoBeads(t)
	client := NewClient()

	tasks, err := client.List()
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}

	t.Logf("Found %d tasks", len(tasks))
	for _, task := range tasks {
		t.Logf("  - %s: %s (status=%s, priority=%d, type=%s)",
			task.ID, task.Title, task.Status, task.Priority, task.Type)
	}
}

func TestClient_ListOpen(t *testing.T) {
	skipIfNoBeads(t)
	client := NewClient()

	tasks, err := client.ListOpen()
	if err != nil {
		t.Fatalf("ListOpen failed: %v", err)
	}

	t.Logf("Found %d open tasks", len(tasks))
	for _, task := range tasks {
		if task.Status != "open" {
			t.Errorf("Expected status 'open', got '%s' for task %s", task.Status, task.ID)
		}
	}
}

func TestClient_Ready(t *testing.T) {
	skipIfNoBeads(t)
	client := NewClient()

	tasks, err := client.Ready()
	if err != nil {
		t.Fatalf("Ready failed: %v", err)
	}

	t.Logf("Found %d ready tasks", len(tasks))
}

func TestClient_Show(t *testing.T) {
	skipIfNoBeads(t)
	client := NewClient()

	// First get a task ID from list
	tasks, err := client.List()
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}

	if len(tasks) == 0 {
		t.Skip("No tasks to show")
	}

	task, err := client.Show(tasks[0].ID)
	if err != nil {
		t.Fatalf("Show failed: %v", err)
	}

	if task.ID != tasks[0].ID {
		t.Errorf("Expected ID %s, got %s", tasks[0].ID, task.ID)
	}

	t.Logf("Showed task: %s - %s", task.ID, task.Title)
}

func TestClient_CreateAndDelete(t *testing.T) {
	skipIfNoBeads(t)
	client := NewClient()

	// Create a test task
	task, err := client.Create(CreateOptions{
		Title:       "Test task from client_test.go",
		Description: "This is a test task",
		Type:        "task",
		Priority:    3,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	t.Logf("Created task: %s - %s", task.ID, task.Title)

	if task.Title != "Test task from client_test.go" {
		t.Errorf("Expected title 'Test task from client_test.go', got '%s'", task.Title)
	}
	if task.Priority != 3 {
		t.Errorf("Expected priority 3, got %d", task.Priority)
	}
	if task.Type != "task" {
		t.Errorf("Expected type 'task', got '%s'", task.Type)
	}

	// Clean up - delete the task
	err = client.Delete(task.ID)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	t.Log("Deleted test task")
}

func TestClient_Update(t *testing.T) {
	skipIfNoBeads(t)
	client := NewClient()

	// Create a test task
	task, err := client.Create(CreateOptions{
		Title:    "Update test task",
		Type:     "task",
		Priority: 2,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	defer client.Delete(task.ID)

	// Update the task
	newPriority := 1
	err = client.Update(task.ID, UpdateOptions{
		Status:   "in_progress",
		Priority: &newPriority,
	})
	if err != nil {
		t.Fatalf("Update failed: %v", err)
	}

	// Verify the update
	updated, err := client.Show(task.ID)
	if err != nil {
		t.Fatalf("Show failed: %v", err)
	}

	if updated.Status != "in_progress" {
		t.Errorf("Expected status 'in_progress', got '%s'", updated.Status)
	}
	if updated.Priority != 1 {
		t.Errorf("Expected priority 1, got %d", updated.Priority)
	}

	t.Log("Update test passed")
}

func TestClient_Close(t *testing.T) {
	skipIfNoBeads(t)
	client := NewClient()

	// Create a test task
	task, err := client.Create(CreateOptions{
		Title:    "Close test task",
		Type:     "task",
		Priority: 3,
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	defer client.Delete(task.ID)

	// Close the task
	err = client.Close(task.ID, "Test completed")
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}

	// Verify the close
	closed, err := client.Show(task.ID)
	if err != nil {
		t.Fatalf("Show failed: %v", err)
	}

	if closed.Status != "closed" {
		t.Errorf("Expected status 'closed', got '%s'", closed.Status)
	}

	t.Log("Close test passed")
}

package alfred

import (
	"encoding/json"
	"testing"

	"azure-workflow/internal/azure"
)

func TestItemHasRequiredFields(t *testing.T) {
	item := Item{
		UID:      "vm",
		Title:    "Virtual Machines",
		Subtitle: "Azure Virtual Machines",
		Arg:      "https://portal.azure.com/...",
	}

	if item.UID != "vm" {
		t.Errorf("expected UID 'vm', got '%s'", item.UID)
	}
	if item.Title != "Virtual Machines" {
		t.Errorf("expected Title 'Virtual Machines', got '%s'", item.Title)
	}
	if item.Subtitle != "Azure Virtual Machines" {
		t.Errorf("expected Subtitle 'Azure Virtual Machines', got '%s'", item.Subtitle)
	}
	if item.Arg != "https://portal.azure.com/..." {
		t.Errorf("expected Arg 'https://portal.azure.com/...', got '%s'", item.Arg)
	}
}

func TestItemHasIconField(t *testing.T) {
	item := Item{
		UID:   "vm",
		Title: "Virtual Machines",
		Icon: Icon{
			Path: "icons/vm.png",
		},
	}

	if item.Icon.Path != "icons/vm.png" {
		t.Errorf("expected Icon.Path 'icons/vm.png', got '%s'", item.Icon.Path)
	}
}

func TestItemCanBeMarshaledToJSON(t *testing.T) {
	item := Item{
		UID:      "vm",
		Title:    "Virtual Machines",
		Subtitle: "Azure Virtual Machines",
		Arg:      "https://portal.azure.com/...",
	}

	_, err := json.Marshal(item)
	if err != nil {
		t.Errorf("failed to marshal Item to JSON: %v", err)
	}
}

func TestItemJSONMatchesAlfredFormat(t *testing.T) {
	item := Item{
		UID:      "vm",
		Title:    "Virtual Machines",
		Subtitle: "Azure Virtual Machines",
		Arg:      "https://portal.azure.com/test",
	}

	jsonBytes, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("failed to marshal Item: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if result["uid"] != "vm" {
		t.Errorf("expected uid 'vm', got '%v'", result["uid"])
	}
	if result["title"] != "Virtual Machines" {
		t.Errorf("expected title 'Virtual Machines', got '%v'", result["title"])
	}
	if result["subtitle"] != "Azure Virtual Machines" {
		t.Errorf("expected subtitle 'Azure Virtual Machines', got '%v'", result["subtitle"])
	}
	if result["arg"] != "https://portal.azure.com/test" {
		t.Errorf("expected arg 'https://portal.azure.com/test', got '%v'", result["arg"])
	}
}

func TestIconFieldMarshaledWithPath(t *testing.T) {
	item := Item{
		UID:   "vm",
		Title: "Virtual Machines",
		Icon: Icon{
			Path: "icons/vm.png",
		},
	}

	jsonBytes, err := json.Marshal(item)
	if err != nil {
		t.Fatalf("failed to marshal Item: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	icon, ok := result["icon"].(map[string]interface{})
	if !ok {
		t.Fatal("expected icon to be an object")
	}
	if icon["path"] != "icons/vm.png" {
		t.Errorf("expected icon.path 'icons/vm.png', got '%v'", icon["path"])
	}
}

// Phase 6: Service를 Alfred Item으로 변환 테스트

func TestServiceCanBeConvertedToItem(t *testing.T) {
	svc := azure.Service{
		Name:    "Virtual Machines",
		Aliases: []string{"vm"},
		URL:     "https://portal.azure.com/vm",
		Icon:    "icons/vm.png",
	}

	item := NewItemFromService(svc)
	if item.Title == "" {
		t.Error("expected item to have a title")
	}
}

func TestItemTitleIsServiceName(t *testing.T) {
	svc := azure.Service{
		Name:    "Virtual Machines",
		Aliases: []string{"vm"},
		URL:     "https://portal.azure.com/vm",
		Icon:    "icons/vm.png",
	}

	item := NewItemFromService(svc)
	if item.Title != "Virtual Machines" {
		t.Errorf("expected Title 'Virtual Machines', got '%s'", item.Title)
	}
}

func TestItemArgIsServiceURL(t *testing.T) {
	svc := azure.Service{
		Name:    "Virtual Machines",
		Aliases: []string{"vm"},
		URL:     "https://portal.azure.com/vm",
		Icon:    "icons/vm.png",
	}

	item := NewItemFromService(svc)
	if item.Arg != "https://portal.azure.com/vm" {
		t.Errorf("expected Arg 'https://portal.azure.com/vm', got '%s'", item.Arg)
	}
}

func TestItemIconPathIsServiceIcon(t *testing.T) {
	svc := azure.Service{
		Name:    "Virtual Machines",
		Aliases: []string{"vm"},
		URL:     "https://portal.azure.com/vm",
		Icon:    "icons/vm.png",
	}

	item := NewItemFromService(svc)
	if item.Icon.Path != "icons/vm.png" {
		t.Errorf("expected Icon.Path 'icons/vm.png', got '%s'", item.Icon.Path)
	}
}

func TestMultipleServicesToItems(t *testing.T) {
	services := []azure.Service{
		{Name: "Virtual Machines", Aliases: []string{"vm"}, URL: "https://portal.azure.com/vm", Icon: "icons/vm.png"},
		{Name: "Storage Accounts", Aliases: []string{"storage"}, URL: "https://portal.azure.com/storage", Icon: "icons/storage.png"},
	}

	items := NewItemsFromServices(services)
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
	if items[0].Title != "Virtual Machines" {
		t.Errorf("expected first item title 'Virtual Machines', got '%s'", items[0].Title)
	}
	if items[1].Title != "Storage Accounts" {
		t.Errorf("expected second item title 'Storage Accounts', got '%s'", items[1].Title)
	}
}

package alfred

import (
	"encoding/json"
	"testing"
)

func TestOutputHasItemsArray(t *testing.T) {
	output := Output{
		Items: []Item{
			{UID: "vm", Title: "Virtual Machines"},
		},
	}

	if len(output.Items) != 1 {
		t.Errorf("expected 1 item, got %d", len(output.Items))
	}
}

func TestOutputJSONContainsItemsKey(t *testing.T) {
	output := Output{
		Items: []Item{
			{UID: "vm", Title: "Virtual Machines"},
		},
	}

	jsonBytes, err := json.Marshal(output)
	if err != nil {
		t.Fatalf("failed to marshal Output: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if _, ok := result["items"]; !ok {
		t.Error("expected 'items' key in JSON output")
	}
}

func TestEmptyItemsReturnsEmptyArray(t *testing.T) {
	output := Output{
		Items: []Item{},
	}

	jsonBytes, err := json.Marshal(output)
	if err != nil {
		t.Fatalf("failed to marshal Output: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	items, ok := result["items"].([]any)
	if !ok {
		t.Fatal("expected items to be an array")
	}
	if len(items) != 0 {
		t.Errorf("expected empty array, got %d items", len(items))
	}
}

func TestMultipleItemsProducesCorrectJSONArray(t *testing.T) {
	output := Output{
		Items: []Item{
			{UID: "vm", Title: "Virtual Machines"},
			{UID: "storage", Title: "Storage Accounts"},
		},
	}

	jsonBytes, err := json.Marshal(output)
	if err != nil {
		t.Fatalf("failed to marshal Output: %v", err)
	}

	var result map[string]any
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	items, ok := result["items"].([]any)
	if !ok {
		t.Fatal("expected items to be an array")
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
}

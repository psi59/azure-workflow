package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"azure-workflow/internal/alfred"
)

func createTestServicesFile(t *testing.T) string {
	content := `services:
  - name: Virtual Machines
    aliases: [vm, virtual machine, 가상머신]
    url: https://portal.azure.com/#blade/HubsExtension/BrowseResourceBlade/resourceType/Microsoft.Compute%2FVirtualMachines
    icon: icons/vm.png
  - name: Storage Accounts
    aliases: [storage, blob, 스토리지]
    url: https://portal.azure.com/#blade/HubsExtension/BrowseResourceBlade/resourceType/Microsoft.Storage%2FStorageAccounts
    icon: icons/storage.png
  - name: AKS
    aliases: [aks, kubernetes, k8s, 쿠버네티스]
    url: https://portal.azure.com/#blade/HubsExtension/BrowseResourceBlade/resourceType/Microsoft.ContainerService%2FmanagedClusters
    icon: icons/aks.png
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "services.yaml")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	return tmpFile
}

func TestProcessQueryReturnsAlfredJSON(t *testing.T) {
	servicesFile := createTestServicesFile(t)
	output, err := ProcessQuery(servicesFile, "vm")
	if err != nil {
		t.Fatalf("ProcessQuery failed: %v", err)
	}

	var result alfred.Output
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		t.Fatalf("failed to parse output as JSON: %v", err)
	}

	if len(result.Items) == 0 {
		t.Error("expected at least one item")
	}
}

func TestStorageQueryReturnsStorageAccounts(t *testing.T) {
	servicesFile := createTestServicesFile(t)
	output, err := ProcessQuery(servicesFile, "storage")
	if err != nil {
		t.Fatalf("ProcessQuery failed: %v", err)
	}

	var result alfred.Output
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		t.Fatalf("failed to parse output as JSON: %v", err)
	}

	if len(result.Items) == 0 {
		t.Fatal("expected at least one item")
	}

	found := false
	for _, item := range result.Items {
		if item.Title == "Storage Accounts" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected 'Storage Accounts' in results")
	}
}

func TestEmptyQueryReturnsAllServices(t *testing.T) {
	servicesFile := createTestServicesFile(t)
	output, err := ProcessQuery(servicesFile, "")
	if err != nil {
		t.Fatalf("ProcessQuery failed: %v", err)
	}

	var result alfred.Output
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		t.Fatalf("failed to parse output as JSON: %v", err)
	}

	if len(result.Items) != 3 {
		t.Errorf("expected 3 items, got %d", len(result.Items))
	}
}

func TestAliasSearchReturnsCorrectResult(t *testing.T) {
	servicesFile := createTestServicesFile(t)
	output, err := ProcessQuery(servicesFile, "k8s")
	if err != nil {
		t.Fatalf("ProcessQuery failed: %v", err)
	}

	var result alfred.Output
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		t.Fatalf("failed to parse output as JSON: %v", err)
	}

	if len(result.Items) == 0 {
		t.Fatal("expected at least one item")
	}

	if result.Items[0].Title != "AKS" {
		t.Errorf("expected first result to be 'AKS', got '%s'", result.Items[0].Title)
	}
}

func TestResultIncludesIconPath(t *testing.T) {
	servicesFile := createTestServicesFile(t)
	output, err := ProcessQuery(servicesFile, "vm")
	if err != nil {
		t.Fatalf("ProcessQuery failed: %v", err)
	}

	// Check that icon path is in the JSON
	if !strings.Contains(output, "icons/vm.png") {
		t.Error("expected output to contain icon path 'icons/vm.png'")
	}
}

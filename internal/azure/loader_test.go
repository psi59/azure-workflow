package azure

import (
	"os"
	"path/filepath"
	"testing"
)

func createTestYAMLFile(t *testing.T) string {
	content := `services:
  - name: Virtual Machines
    aliases: [vm, virtual machine, 가상머신]
    url: https://portal.azure.com/#blade/HubsExtension/BrowseResourceBlade/resourceType/Microsoft.Compute%2FVirtualMachines
    icon: icons/vm.png
`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "services.yaml")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	return tmpFile
}

func TestLoadServicesFromYAML(t *testing.T) {
	tmpFile := createTestYAMLFile(t)

	services, err := LoadServices(tmpFile)
	if err != nil {
		t.Fatalf("LoadServices failed: %v", err)
	}

	if len(services) == 0 {
		t.Error("expected at least one service, got none")
	}
}

func TestLoadedServiceHasName(t *testing.T) {
	tmpFile := createTestYAMLFile(t)

	services, err := LoadServices(tmpFile)
	if err != nil {
		t.Fatalf("LoadServices failed: %v", err)
	}

	if services[0].Name != "Virtual Machines" {
		t.Errorf("expected Name 'Virtual Machines', got '%s'", services[0].Name)
	}
}

func TestLoadedServiceHasAliases(t *testing.T) {
	tmpFile := createTestYAMLFile(t)

	services, err := LoadServices(tmpFile)
	if err != nil {
		t.Fatalf("LoadServices failed: %v", err)
	}

	expectedAliases := []string{"vm", "virtual machine", "가상머신"}
	if len(services[0].Aliases) != len(expectedAliases) {
		t.Errorf("expected %d aliases, got %d", len(expectedAliases), len(services[0].Aliases))
	}
	for i, expected := range expectedAliases {
		if services[0].Aliases[i] != expected {
			t.Errorf("expected alias[%d] '%s', got '%s'", i, expected, services[0].Aliases[i])
		}
	}
}

func TestLoadedServiceHasURL(t *testing.T) {
	tmpFile := createTestYAMLFile(t)

	services, err := LoadServices(tmpFile)
	if err != nil {
		t.Fatalf("LoadServices failed: %v", err)
	}

	expectedURL := "https://portal.azure.com/#blade/HubsExtension/BrowseResourceBlade/resourceType/Microsoft.Compute%2FVirtualMachines"
	if services[0].URL != expectedURL {
		t.Errorf("expected URL '%s', got '%s'", expectedURL, services[0].URL)
	}
}

func TestLoadedServiceHasIcon(t *testing.T) {
	tmpFile := createTestYAMLFile(t)

	services, err := LoadServices(tmpFile)
	if err != nil {
		t.Fatalf("LoadServices failed: %v", err)
	}

	if services[0].Icon != "icons/vm.png" {
		t.Errorf("expected Icon 'icons/vm.png', got '%s'", services[0].Icon)
	}
}

func TestLoadServicesReturnsErrorForNonExistentFile(t *testing.T) {
	_, err := LoadServices("/non/existent/path/services.yaml")
	if err == nil {
		t.Error("expected error for non-existent file, got nil")
	}
}

func TestLoadServicesReturnsErrorForInvalidYAML(t *testing.T) {
	content := `invalid yaml: [unclosed bracket`
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "invalid.yaml")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	_, err := LoadServices(tmpFile)
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}

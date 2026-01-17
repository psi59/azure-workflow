package azure

import (
	"testing"
)

func TestServiceHasRequiredFields(t *testing.T) {
	s := Service{
		Name:    "Virtual Machines",
		Aliases: []string{"vm", "virtual machine"},
		URL:     "https://portal.azure.com/#blade/HubsExtension/BrowseResourceBlade/resourceType/Microsoft.Compute%2FVirtualMachines",
		Icon:    "icons/vm.png",
	}

	if s.Name != "Virtual Machines" {
		t.Errorf("expected Name to be 'Virtual Machines', got '%s'", s.Name)
	}
	if len(s.Aliases) != 2 {
		t.Errorf("expected 2 aliases, got %d", len(s.Aliases))
	}
	if s.Aliases[0] != "vm" {
		t.Errorf("expected first alias to be 'vm', got '%s'", s.Aliases[0])
	}
	if s.URL == "" {
		t.Error("expected URL to be non-empty")
	}
	if s.Icon != "icons/vm.png" {
		t.Errorf("expected Icon to be 'icons/vm.png', got '%s'", s.Icon)
	}
}

func TestServiceAliasesContainsShortNamesAndKorean(t *testing.T) {
	s := Service{
		Name:    "AKS",
		Aliases: []string{"aks", "kubernetes", "k8s", "쿠버네티스"},
		URL:     "https://portal.azure.com/#blade/HubsExtension/BrowseResourceBlade/resourceType/Microsoft.ContainerService%2FmanagedClusters",
		Icon:    "icons/aks.png",
	}

	expectedAliases := []string{"aks", "kubernetes", "k8s", "쿠버네티스"}
	if len(s.Aliases) != len(expectedAliases) {
		t.Errorf("expected %d aliases, got %d", len(expectedAliases), len(s.Aliases))
	}

	for i, expected := range expectedAliases {
		if s.Aliases[i] != expected {
			t.Errorf("expected alias[%d] to be '%s', got '%s'", i, expected, s.Aliases[i])
		}
	}
}

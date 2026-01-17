package search

import (
	"testing"

	"azure-workflow/internal/azure"
)

func getTestServices() []azure.Service {
	return []azure.Service{
		{Name: "Virtual Machines", Aliases: []string{"vm", "virtual machine", "가상머신"}, URL: "https://portal.azure.com/vm", Icon: "icons/vm.png"},
		{Name: "Storage Accounts", Aliases: []string{"storage", "blob", "스토리지"}, URL: "https://portal.azure.com/storage", Icon: "icons/storage.png"},
		{Name: "AKS", Aliases: []string{"aks", "kubernetes", "k8s", "쿠버네티스"}, URL: "https://portal.azure.com/aks", Icon: "icons/aks.png"},
		{Name: "Azure OpenAI Service", Aliases: []string{"openai", "gpt", "chatgpt", "ai", "llm"}, URL: "https://portal.azure.com/openai", Icon: "icons/openai.png"},
	}
}

func TestSearchByExactName(t *testing.T) {
	services := getTestServices()
	results := Search(services, "Virtual Machines")

	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}
	if results[0].Name != "Virtual Machines" {
		t.Errorf("expected 'Virtual Machines', got '%s'", results[0].Name)
	}
}

func TestSearchCaseInsensitive(t *testing.T) {
	services := getTestServices()
	results := Search(services, "virtual machines")

	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}
	if results[0].Name != "Virtual Machines" {
		t.Errorf("expected 'Virtual Machines', got '%s'", results[0].Name)
	}
}

func TestSearchBySubstring(t *testing.T) {
	services := getTestServices()
	results := Search(services, "Storage")

	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}
	if results[0].Name != "Storage Accounts" {
		t.Errorf("expected 'Storage Accounts', got '%s'", results[0].Name)
	}
}

func TestSearchByAlias(t *testing.T) {
	services := getTestServices()
	results := Search(services, "vm")

	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}
	if results[0].Name != "Virtual Machines" {
		t.Errorf("expected 'Virtual Machines', got '%s'", results[0].Name)
	}
}

func TestSearchByShortName(t *testing.T) {
	services := getTestServices()
	results := Search(services, "k8s")

	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}
	if results[0].Name != "AKS" {
		t.Errorf("expected 'AKS', got '%s'", results[0].Name)
	}
}

func TestSearchResultsSortedByRelevance(t *testing.T) {
	services := getTestServices()
	results := Search(services, "aks")

	if len(results) == 0 {
		t.Fatal("expected at least one result")
	}
	// AKS should be first as it's an exact alias match
	if results[0].Name != "AKS" {
		t.Errorf("expected 'AKS' as first result, got '%s'", results[0].Name)
	}
}

func TestEmptyQueryReturnsAllServices(t *testing.T) {
	services := getTestServices()
	results := Search(services, "")

	if len(results) != len(services) {
		t.Errorf("expected %d services, got %d", len(services), len(results))
	}
}

func TestNoMatchReturnsEmptyResult(t *testing.T) {
	services := getTestServices()
	results := Search(services, "xyznonexistent")

	if len(results) != 0 {
		t.Errorf("expected 0 results, got %d", len(results))
	}
}

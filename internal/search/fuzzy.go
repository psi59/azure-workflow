package search

import (
	"strings"

	"azure-workflow/internal/azure"
	"github.com/sahilm/fuzzy"
)

type searchable struct {
	services []azure.Service
	targets  []string
}

func (s searchable) String(i int) string {
	return s.targets[i]
}

func (s searchable) Len() int {
	return len(s.targets)
}

func Search(services []azure.Service, query string) []azure.Service {
	if query == "" {
		return services
	}

	// 각 서비스에 대해 이름과 별칭을 합친 검색 대상 생성
	targets := make([]string, len(services))
	for i, svc := range services {
		parts := []string{svc.Name}
		parts = append(parts, svc.Aliases...)
		targets[i] = strings.ToLower(strings.Join(parts, " "))
	}

	source := searchable{
		services: services,
		targets:  targets,
	}

	matches := fuzzy.FindFrom(strings.ToLower(query), source)

	results := make([]azure.Service, len(matches))
	for i, match := range matches {
		results[i] = services[match.Index]
	}

	return results
}

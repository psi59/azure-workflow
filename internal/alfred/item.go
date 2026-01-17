package alfred

import (
	"strings"

	"azure-workflow/internal/azure"
)

type Icon struct {
	Path string `json:"path"`
}

type Item struct {
	UID      string `json:"uid"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Arg      string `json:"arg"`
	Icon     Icon   `json:"icon"`
}

func NewItemFromService(svc azure.Service) Item {
	return Item{
		UID:      strings.ToLower(strings.ReplaceAll(svc.Name, " ", "-")),
		Title:    svc.Name,
		Subtitle: strings.Join(svc.Aliases, ", "),
		Arg:      svc.URL,
		Icon:     Icon{Path: svc.Icon},
	}
}

func NewItemsFromServices(services []azure.Service) []Item {
	items := make([]Item, len(services))
	for i, svc := range services {
		items[i] = NewItemFromService(svc)
	}
	return items
}

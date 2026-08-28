package content

import (
	"fmt"
	"sort"
	"strings"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type ReferenceManager struct{ repo *store.Repository }

func NewReferenceManager(repo *store.Repository) *ReferenceManager {
	return &ReferenceManager{repo: repo}
}

func (m *ReferenceManager) Save(reference domain.CaseReference) error {
	reference.Label = strings.TrimSpace(reference.Label)
	reference.Source = strings.TrimSpace(reference.Source)
	reference.Location = strings.TrimSpace(reference.Location)
	if err := reference.Validate(); err != nil {
		return err
	}
	return m.repo.PutReference(reference)
}

func (m *ReferenceManager) List(caseID string) ([]domain.CaseReference, error) {
	items, err := m.repo.ListReferences(caseID)
	if err != nil {
		return nil, err
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Source != items[j].Source {
			return items[i].Source < items[j].Source
		}
		return items[i].Location < items[j].Location
	})
	return items, nil
}

func (m *ReferenceManager) Citation(caseID string) (string, error) {
	items, err := m.List(caseID)
	if err != nil {
		return "", err
	}
	if len(items) == 0 {
		return "", fmt.Errorf("case %s has no references", caseID)
	}
	parts := make([]string, 0, len(items))
	for _, item := range items {
		location := item.Location
		if location == "" {
			location = "未标注位置"
		}
		parts = append(parts, fmt.Sprintf("%s (%s, %s)", item.Label, item.Source, location))
	}
	return strings.Join(parts, "; "), nil
}

func ReferenceSources(items []domain.CaseReference) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(items))
	for _, item := range items {
		if !seen[item.Source] {
			seen[item.Source] = true
			result = append(result, item.Source)
		}
	}
	sort.Strings(result)
	return result
}

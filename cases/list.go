package cases

import (
	"sort"

	"example.com/casescript/domain"
)

func (m *Manager) ListPublishedByDate(date string) ([]domain.LegalCase, error) {
	items, err := m.repo.ListCases()
	if err != nil {
		return nil, err
	}
	filtered := make([]domain.LegalCase, 0, len(items))
	for _, value := range items {
		if value.PublishDate == date && value.IsPublished() {
			filtered = append(filtered, value)
		}
	}
	sort.SliceStable(filtered, func(i, j int) bool {
		return filtered[i].PublishDate < filtered[j].PublishDate
	})
	return filtered, nil
}

func (m *Manager) ListAll() ([]domain.LegalCase, error) { return m.repo.ListCases() }

func GroupByDate(items []domain.LegalCase) map[string][]domain.LegalCase {
	groups := make(map[string][]domain.LegalCase)
	for _, value := range items {
		groups[value.PublishDate] = append(groups[value.PublishDate], value)
	}
	return groups
}

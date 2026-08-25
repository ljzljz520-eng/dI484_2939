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
	// Cases sharing a publish date all share the same PublishDate, so comparing on
	// PublishDate alone leaves the underlying (bucket key / insertion) order intact.
	// That order is not stable across runs, which makes per-date classroom links point
	// at different content each load. Tie-break on the normalized case number to match
	// ApplyFilter and BuildTeacherBoard, then on the stable case id as a final guard.
	sort.SliceStable(filtered, func(i, j int) bool {
		if filtered[i].PublishDate != filtered[j].PublishDate {
			return filtered[i].PublishDate < filtered[j].PublishDate
		}
		if keyI, keyJ := domain.CaseNumberKey(filtered[i].Number), domain.CaseNumberKey(filtered[j].Number); keyI != keyJ {
			return keyI < keyJ
		}
		return filtered[i].ID < filtered[j].ID
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

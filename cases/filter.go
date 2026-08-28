package cases

import (
	"sort"
	"strings"

	"example.com/casescript/domain"
)

type CaseFilter struct {
	Date          string
	Status        domain.CaseStatus
	Text          string
	RequireNumber string
}

func ApplyFilter(items []domain.LegalCase, filter CaseFilter) []domain.LegalCase {
	result := make([]domain.LegalCase, 0, len(items))
	needle := strings.ToLower(strings.TrimSpace(filter.Text))
	for _, value := range items {
		if filter.Date != "" && value.PublishDate != filter.Date {
			continue
		}
		if filter.Status != "" && value.Status != filter.Status {
			continue
		}
		if filter.RequireNumber != "" && !strings.HasPrefix(domain.CaseNumberKey(value.Number), domain.CaseNumberKey(filter.RequireNumber)) {
			continue
		}
		if needle != "" && !strings.Contains(strings.ToLower(value.Title), needle) && !strings.Contains(strings.ToLower(value.Summary), needle) {
			continue
		}
		result = append(result, value)
	}
	sort.SliceStable(result, func(i, j int) bool {
		if result[i].PublishDate != result[j].PublishDate {
			return result[i].PublishDate < result[j].PublishDate
		}
		return domain.CaseNumberKey(result[i].Number) < domain.CaseNumberKey(result[j].Number)
	})
	return result
}

func FilterDescription(filter CaseFilter) string {
	parts := make([]string, 0, 4)
	if filter.Date != "" {
		parts = append(parts, "date="+filter.Date)
	}
	if filter.Status != "" {
		parts = append(parts, "status="+string(filter.Status))
	}
	if filter.Text != "" {
		parts = append(parts, "text="+strings.TrimSpace(filter.Text))
	}
	if filter.RequireNumber != "" {
		parts = append(parts, "number="+filter.RequireNumber)
	}
	return strings.Join(parts, ", ")
}

package cases

import (
	"strings"

	"example.com/casescript/domain"
)

func Search(items []domain.LegalCase, term string) []domain.LegalCase {
	needle := strings.ToLower(strings.TrimSpace(term))
	if needle == "" {
		return append([]domain.LegalCase(nil), items...)
	}
	result := make([]domain.LegalCase, 0)
	for _, value := range items {
		if strings.Contains(strings.ToLower(value.Title), needle) || strings.Contains(strings.ToLower(value.Summary), needle) {
			result = append(result, value)
		}
	}
	return result
}

func PublishedOnly(items []domain.LegalCase) []domain.LegalCase {
	result := make([]domain.LegalCase, 0, len(items))
	for _, value := range items {
		if value.Status == domain.StatusPublished {
			result = append(result, value)
		}
	}
	return result
}

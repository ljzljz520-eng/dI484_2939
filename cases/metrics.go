package cases

import (
	"sort"

	"example.com/casescript/domain"
)

type DateMetric struct {
	Date        string
	Total       int
	Published   int
	Drafts      int
	Archived    int
	CaseNumbers []string
}

func BuildDateMetrics(items []domain.LegalCase) []DateMetric {
	byDate := make(map[string]*DateMetric)
	for _, value := range items {
		metric := byDate[value.PublishDate]
		if metric == nil {
			metric = &DateMetric{Date: value.PublishDate, CaseNumbers: make([]string, 0)}
			byDate[value.PublishDate] = metric
		}
		metric.Total++
		metric.CaseNumbers = append(metric.CaseNumbers, value.Number)
		switch value.Status {
		case domain.StatusPublished:
			metric.Published++
		case domain.StatusDraft:
			metric.Drafts++
		case domain.StatusArchived:
			metric.Archived++
		}
	}
	result := make([]DateMetric, 0, len(byDate))
	for _, metric := range byDate {
		sort.Strings(metric.CaseNumbers)
		result = append(result, *metric)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Date < result[j].Date })
	return result
}

func (m DateMetric) ReadyForTeaching() bool { return m.Published > 0 && m.Published == m.Total }

func (m DateMetric) CoverageRatio() float64 {
	if m.Total == 0 {
		return 0
	}
	return float64(m.Published) / float64(m.Total)
}

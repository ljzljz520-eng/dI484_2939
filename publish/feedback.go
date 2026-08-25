package publish

import (
	"sort"
	"strings"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type StudentFeedback = domain.StudentFeedback

type FeedbackManager struct{ repo *store.Repository }

func NewFeedbackManager(repo *store.Repository) *FeedbackManager { return &FeedbackManager{repo: repo} }

func (m *FeedbackManager) Save(value StudentFeedback) error {
	value.Student = strings.TrimSpace(value.Student)
	value.Text = strings.TrimSpace(value.Text)
	if err := value.Validate(); err != nil {
		return err
	}
	return m.repo.PutFeedback(value)
}

func (m *FeedbackManager) List(caseID string) ([]StudentFeedback, error) {
	items, err := m.repo.ListFeedback(caseID)
	if err != nil {
		return nil, err
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Date != items[j].Date {
			return items[i].Date < items[j].Date
		}
		return items[i].ID < items[j].ID
	})
	return items, nil
}

func (m *FeedbackManager) Resolve(id string) error {
	value, err := m.repo.GetFeedback(id)
	if err != nil {
		return err
	}
	value.Resolved = true
	return m.Save(value)
}

func OpenFeedbackCount(items []StudentFeedback) int {
	count := 0
	for _, item := range items {
		if !item.Resolved {
			count++
		}
	}
	return count
}

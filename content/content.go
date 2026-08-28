package content

import (
	"fmt"
	"sort"
	"strings"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type Manager struct{ repo *store.Repository }

func NewManager(repo *store.Repository) *Manager { return &Manager{repo: repo} }

func (m *Manager) AddChapter(value domain.Chapter) error {
	value.Title = strings.TrimSpace(value.Title)
	value.Body = strings.TrimSpace(value.Body)
	if err := value.Validate(); err != nil {
		return err
	}
	return m.repo.PutChapter(value)
}

func (m *Manager) Chapters(caseID string) ([]domain.Chapter, error) {
	items, err := m.repo.ListChapters(caseID)
	if err != nil {
		return nil, err
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].Position == items[j].Position {
			return items[i].ID < items[j].ID
		}
		return items[i].Position < items[j].Position
	})
	return items, nil
}

func (m *Manager) SaveNote(value domain.ClassroomNote) error {
	value.Text = strings.TrimSpace(value.Text)
	value.Author = strings.TrimSpace(value.Author)
	if err := value.Validate(); err != nil {
		return err
	}
	return m.repo.PutNote(value)
}

func (m *Manager) Notes(caseID string) ([]domain.ClassroomNote, error) {
	items, err := m.repo.ListNotes(caseID)
	if err != nil {
		return nil, err
	}
	sort.Slice(items, func(i, j int) bool { return items[i].ID < items[j].ID })
	return items, nil
}

func (m *Manager) ValidateDraft(caseID string) error {
	chapters, err := m.Chapters(caseID)
	if err != nil {
		return err
	}
	if len(chapters) == 0 {
		return fmt.Errorf("case %s requires at least one chapter", caseID)
	}
	for _, chapter := range chapters {
		if strings.TrimSpace(chapter.Body) == "" {
			return fmt.Errorf("chapter %s has empty body", chapter.ID)
		}
	}
	return nil
}

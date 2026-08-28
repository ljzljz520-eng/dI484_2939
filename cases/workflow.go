package cases

import (
	"fmt"
	"sort"
	"strings"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type CaseWorkspace struct {
	Case       domain.LegalCase
	Chapters   []domain.Chapter
	Notes      []domain.ClassroomNote
	References []domain.CaseReference
	Reviews    []domain.ReviewRequest
}

func (m *Manager) LoadWorkspace(caseID string) (CaseWorkspace, error) {
	value, err := m.repo.GetCase(caseID)
	if err != nil {
		return CaseWorkspace{}, err
	}
	chapters, err := m.repo.ListChapters(caseID)
	if err != nil {
		return CaseWorkspace{}, err
	}
	notes, err := m.repo.ListNotes(caseID)
	if err != nil {
		return CaseWorkspace{}, err
	}
	references, err := m.repo.ListReferences(caseID)
	if err != nil {
		return CaseWorkspace{}, err
	}
	reviews, err := m.repo.ListReviews(caseID)
	if err != nil {
		return CaseWorkspace{}, err
	}
	sort.Slice(chapters, func(i, j int) bool { return chapters[i].Position < chapters[j].Position })
	sort.Slice(notes, func(i, j int) bool { return notes[i].UpdatedAt < notes[j].UpdatedAt })
	return CaseWorkspace{Case: value, Chapters: chapters, Notes: notes, References: references, Reviews: reviews}, nil
}

func (w CaseWorkspace) Validate() error {
	if err := w.Case.Validate(); err != nil {
		return err
	}
	if len(w.Chapters) == 0 {
		return fmt.Errorf("workspace requires at least one chapter")
	}
	for _, chapter := range w.Chapters {
		if chapter.CaseID != w.Case.ID {
			return fmt.Errorf("chapter %s belongs to another case", chapter.ID)
		}
	}
	for _, note := range w.Notes {
		if note.CaseID != w.Case.ID {
			return fmt.Errorf("note %s belongs to another case", note.ID)
		}
	}
	return nil
}

func (w CaseWorkspace) Search(term string) CaseWorkspace {
	term = strings.ToLower(strings.TrimSpace(term))
	if term == "" {
		return w
	}
	filtered := make([]domain.Chapter, 0, len(w.Chapters))
	for _, chapter := range w.Chapters {
		if strings.Contains(strings.ToLower(chapter.Title), term) || strings.Contains(strings.ToLower(chapter.Body), term) {
			filtered = append(filtered, chapter)
		}
	}
	w.Chapters = filtered
	return w
}

func WorkspaceSummary(w CaseWorkspace) domain.CaseSummary {
	return domain.CaseSummary{CaseID: w.Case.ID, ChapterCount: len(w.Chapters), NoteCount: len(w.Notes), ReferenceCount: len(w.References), IsPublished: w.Case.IsPublished()}
}

func NewWorkspaceManager(repo *store.Repository) *Manager { return NewManager(repo) }

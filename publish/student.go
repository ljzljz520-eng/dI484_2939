package publish

import (
	"fmt"
	"sort"

	"example.com/casescript/cases"
	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type StudentView struct {
	Case      domain.LegalCase
	Chapters  []domain.Chapter
	Citations []domain.CaseReference
	Summary   string
}

func BuildStudentView(repo *store.Repository, caseID string) (StudentView, error) {
	caseValue, err := repo.GetCase(caseID)
	if err != nil {
		return StudentView{}, err
	}
	if !IsVisibleToStudent(caseValue) {
		return StudentView{}, fmt.Errorf("case %s is not visible to students", caseID)
	}
	chapters, err := repo.ListChapters(caseID)
	if err != nil {
		return StudentView{}, err
	}
	references, err := repo.ListReferences(caseID)
	if err != nil {
		return StudentView{}, err
	}
	sort.Slice(chapters, func(i, j int) bool {
		if chapters[i].Position != chapters[j].Position {
			return chapters[i].Position < chapters[j].Position
		}
		return chapters[i].ID < chapters[j].ID
	})
	sort.Slice(references, func(i, j int) bool { return references[i].Label < references[j].Label })
	return StudentView{Case: caseValue, Chapters: chapters, Citations: references, Summary: cases.DraftLabel(caseValue)}, nil
}

func StudentChapterTitles(view StudentView) []string {
	result := make([]string, 0, len(view.Chapters))
	for _, chapter := range view.Chapters {
		result = append(result, chapter.Title)
	}
	return result
}

func StudentCanOpen(view StudentView) bool { return view.Case.IsPublished() && len(view.Chapters) > 0 }

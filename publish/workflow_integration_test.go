package publish

import (
	"path/filepath"
	"strings"
	"testing"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

func TestWorkflowPublishAndExportHandout(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "cases.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()
	if err := repo.PutCase(domain.LegalCase{ID: "case-1", Number: "2024-001", Title: "合同", Summary: "课堂摘要", PublishDate: "2024-06-01", Status: domain.StatusDraft}); err != nil {
		t.Fatal(err)
	}
	if err := repo.PutChapter(domain.Chapter{ID: "chapter-1", CaseID: "case-1", Title: "争点", Body: "争点内容", Position: 1}); err != nil {
		t.Fatal(err)
	}
	if err := repo.PutNote(domain.ClassroomNote{ID: "note-1", CaseID: "case-1", ChapterID: "chapter-1", Text: "讲解提示", Author: "教师", UpdatedAt: "2024-06-01"}); err != nil {
		t.Fatal(err)
	}
	service := NewService(repo)
	publication, err := service.PublishCase("case-1", "2024-06-02")
	if err != nil || publication.Version != 1 {
		t.Fatalf("publication=%#v err=%v", publication, err)
	}
	handout, err := service.ExportHandout("case-1")
	if err != nil || !strings.Contains(handout, "讲解提示") || !strings.Contains(handout, "争点") {
		t.Fatalf("handout=%q err=%v", handout, err)
	}
	sections := HandoutSections(handout)
	if len(sections) < 3 {
		t.Fatalf("sections=%#v", sections)
	}
}

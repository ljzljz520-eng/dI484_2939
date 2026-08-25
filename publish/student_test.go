package publish

import (
	"path/filepath"
	"testing"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

func TestStudentViewOnlyPublished(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "cases.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()
	value := domain.LegalCase{ID: "case-1", Number: "2024-001", Title: "已发布", Summary: "摘要", PublishDate: "2024-06-01", Status: domain.StatusPublished}
	if err := repo.PutCase(value); err != nil {
		t.Fatal(err)
	}
	if err := repo.PutChapter(domain.Chapter{ID: "ch-1", CaseID: value.ID, Title: "第一章", Body: "内容", Position: 1}); err != nil {
		t.Fatal(err)
	}
	view, err := BuildStudentView(repo, value.ID)
	if err != nil || !StudentCanOpen(view) || len(StudentChapterTitles(view)) != 1 {
		t.Fatalf("view=%#v err=%v", view, err)
	}
}

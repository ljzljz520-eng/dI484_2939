package store

import (
	"path/filepath"
	"testing"

	"example.com/casescript/domain"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	path := filepath.Join(t.TempDir(), "cases.db")
	repo, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	items := []struct {
		name string
		put  func() error
	}{
		{"case", func() error {
			return repo.PutCase(domain.LegalCase{ID: "case-1", Number: "2024-001", Title: "标题", Summary: "摘要", PublishDate: "2024-06-01", Status: domain.StatusPublished})
		}},
		{"chapter", func() error {
			return repo.PutChapter(domain.Chapter{ID: "chapter-1", CaseID: "case-1", Title: "事实", Body: "内容", Position: 1})
		}},
		{"note", func() error {
			return repo.PutNote(domain.ClassroomNote{ID: "note-1", CaseID: "case-1", Text: "提示", Author: "教师", UpdatedAt: "2024-06-01"})
		}},
		{"publication", func() error {
			return repo.PutPublication(domain.Publication{CaseID: "case-1", Status: domain.StatusPublished, PublishedAt: "2024-06-01", Version: 1})
		}},
	}
	for _, item := range items {
		if err := item.put(); err != nil {
			t.Fatalf("%s: %v", item.name, err)
		}
	}
	if err := repo.Close(); err != nil {
		t.Fatal(err)
	}
	reopened, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer reopened.Close()
	if value, err := reopened.GetCase("case-1"); err != nil || value.Title != "标题" {
		t.Fatalf("case=%#v err=%v", value, err)
	}
	if chapters, err := reopened.ListChapters("case-1"); err != nil || len(chapters) != 1 {
		t.Fatalf("chapters=%#v err=%v", chapters, err)
	}
	if notes, err := reopened.ListNotes("case-1"); err != nil || len(notes) != 1 {
		t.Fatalf("notes=%#v err=%v", notes, err)
	}
	if publication, err := reopened.GetPublication("case-1"); err != nil || publication.Version != 1 {
		t.Fatalf("publication=%#v err=%v", publication, err)
	}
}

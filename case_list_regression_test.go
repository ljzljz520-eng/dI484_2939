package casescript

import (
	"path/filepath"
	"testing"

	"example.com/casescript/cases"
	"example.com/casescript/domain"
	"example.com/casescript/store"
)

func TestCaseListStableForSameDate(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "cases.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()
	manager := cases.NewManager(repo)
	values := []domain.LegalCase{
		{ID: "case-z", Number: "2024-001", Title: "第一案", Summary: "同日资料一", PublishDate: "2024-06-01", Status: domain.StatusPublished, CreatedAt: "2024-05-01"},
		{ID: "case-a", Number: "2024-002", Title: "第二案", Summary: "同日资料二", PublishDate: "2024-06-01", Status: domain.StatusPublished, CreatedAt: "2024-05-01"},
	}
	for _, value := range values {
		if err := repo.PutCase(value); err != nil {
			t.Fatal(err)
		}
	}
	items, err := manager.ListPublishedByDate("2024-06-01")
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 2 || items[0].Number != "2024-001" || items[1].Number != "2024-002" {
		t.Fatalf("unexpected stable order: %#v", items)
	}
}

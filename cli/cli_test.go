package cli

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

func TestCLIListAndSearch(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "cases.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()
	if err := repo.PutCase(domain.LegalCase{ID: "case-1", Number: "2024-001", Title: "合同", Summary: "摘要", PublishDate: "2024-06-01", Status: domain.StatusPublished}); err != nil {
		t.Fatal(err)
	}
	runner := New(repo)
	var output bytes.Buffer
	if err := runner.Run([]string{"list", "-date", "2024-06-01"}, &output); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "2024-001") {
		t.Fatal(output.String())
	}
	output.Reset()
	if err := runner.Run([]string{"search", "-text", "合同"}, &output); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(output.String(), "合同") {
		t.Fatal(output.String())
	}
}

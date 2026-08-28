package publish

import (
	"path/filepath"
	"testing"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

func TestFeedbackLifecycle(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "cases.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()
	manager := NewFeedbackManager(repo)
	if err := manager.Save(StudentFeedback{ID: "feedback-1", CaseID: "case-1", Student: "学生甲", Text: "希望补充时间线", Date: "2024-06-01"}); err != nil {
		t.Fatal(err)
	}
	items, err := manager.List("case-1")
	if err != nil || len(items) != 1 || OpenFeedbackCount(items) != 1 {
		t.Fatalf("items=%#v err=%v", items, err)
	}
	if err := manager.Resolve("feedback-1"); err != nil {
		t.Fatal(err)
	}
	items, err = manager.List("case-1")
	if err != nil || OpenFeedbackCount(items) != 0 || items[0].Resolved != true {
		t.Fatalf("resolved items=%#v err=%v", items, err)
	}
	if !domain.CanRead(domain.RoleStudent, domain.LegalCase{Status: domain.StatusPublished}) {
		t.Fatal("student access denied")
	}
}

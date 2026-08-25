package cases

import (
	"path/filepath"
	"testing"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

func TestWorkflowCaseIntakeAndWorkspace(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "cases.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()
	manager := NewManager(repo)
	value, err := manager.CreateFromDraft(CreateInput{ID: "case-1", Number: "2024-001", Title: "合同解释", Summary: "课堂摘要", PublishDate: "2024-06-01", CreatedAt: "2024-05-01"}, IntakeChecklist{HasSourceMaterial: true, HasCourtNumber: true, HasLearningGoal: true, HasPublishDate: true})
	if err != nil {
		t.Fatal(err)
	}
	if value.Status != domain.StatusDraft {
		t.Fatalf("status=%s", value.Status)
	}
	chapters := []domain.Chapter{{ID: "ch-2", CaseID: value.ID, Title: "争点", Body: "争点内容", Position: 2}, {ID: "ch-1", CaseID: value.ID, Title: "事实", Body: "事实内容", Position: 1}}
	for _, chapter := range chapters {
		if err := repo.PutChapter(chapter); err != nil {
			t.Fatal(err)
		}
	}
	workspace, err := manager.LoadWorkspace(value.ID)
	if err != nil {
		t.Fatal(err)
	}
	if err := workspace.Validate(); err != nil {
		t.Fatal(err)
	}
	if workspace.Chapters[0].Position != 1 || casesSummary(workspace).ChapterCount != 2 {
		t.Fatalf("workspace=%#v", workspace)
	}
}

func casesSummary(workspace CaseWorkspace) domain.CaseSummary { return WorkspaceSummary(workspace) }

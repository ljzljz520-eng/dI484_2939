package content

import (
	"path/filepath"
	"testing"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

func TestWorkflowTeacherBuildsDraftMaterials(t *testing.T) {
	repo, err := store.Open(filepath.Join(t.TempDir(), "cases.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()
	planManager := NewPlanManager(repo)
	if err := planManager.SavePlan(domain.LessonPlan{ID: "plan-1", CaseID: "case-1", LearningGoal: "识别争点", Audience: domain.AudienceStudent, DurationMin: 45, Activities: []string{"导入", "讨论"}}); err != nil {
		t.Fatal(err)
	}
	if _, err := planManager.AddActivity("case-1", "总结"); err != nil {
		t.Fatal(err)
	}
	refs := NewReferenceManager(repo)
	if err := refs.Save(domain.CaseReference{ID: "ref-1", CaseID: "case-1", Label: "判决书", Source: "最高院", Location: "第3页"}); err != nil {
		t.Fatal(err)
	}
	citation, err := refs.Citation("case-1")
	if err != nil || citation == "" {
		t.Fatalf("citation=%q err=%v", citation, err)
	}
	review := NewReviewManager(repo)
	request, err := review.Request("case-1", "教师甲", "2024-06-01")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := review.Decide(request.ID, "approved", "可发布"); err != nil {
		t.Fatal(err)
	}
	pending, err := review.Pending("case-1")
	if err != nil || len(pending) != 0 {
		t.Fatalf("pending=%#v err=%v", pending, err)
	}
}

package store

import (
	"path/filepath"
	"testing"

	"example.com/casescript/domain"
)

func TestStoreAdditionalEntities(t *testing.T) {
	repo, err := Open(filepath.Join(t.TempDir(), "cases.db"))
	if err != nil {
		t.Fatal(err)
	}
	defer repo.Close()
	if err := repo.PutLessonPlan(domain.LessonPlan{ID: "plan-1", CaseID: "case-1", LearningGoal: "理解举证", Audience: domain.AudienceStudent, DurationMin: 45, Activities: []string{"阅读", "讨论"}}); err != nil {
		t.Fatal(err)
	}
	if err := repo.PutReference(domain.CaseReference{ID: "ref-1", CaseID: "case-1", Label: "法条", Source: "民法典", Location: "第十条"}); err != nil {
		t.Fatal(err)
	}
	if err := repo.PutReview(domain.ReviewRequest{ID: "review-1", CaseID: "case-1", Reviewer: "教师", RequestedAt: "2024-06-01"}); err != nil {
		t.Fatal(err)
	}
	if err := repo.PutTimelineEvent(domain.TimelineEvent{ID: "event-1", CaseID: "case-1", Date: "2024-06-01", Kind: "created", Actor: "教师", Message: "录入", Sequence: 1}); err != nil {
		t.Fatal(err)
	}
	if plans, _ := repo.ListLessonPlans("case-1"); len(plans) != 1 {
		t.Fatal("lesson plan missing")
	}
	if refs, _ := repo.ListReferences("case-1"); len(refs) != 1 {
		t.Fatal("reference missing")
	}
	if reviews, _ := repo.ListReviews("case-1"); len(reviews) != 1 {
		t.Fatal("review missing")
	}
	if events, _ := repo.ListTimeline("case-1"); len(events) != 1 {
		t.Fatal("timeline missing")
	}
}

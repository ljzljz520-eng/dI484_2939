package content

import "testing"

func TestAssessmentOrderingAndRubric(t *testing.T) {
	prompts := []DiscussionPrompt{{ID: "p-2", CaseID: "case-1", ChapterID: "ch-1", Prompt: "第二问", Expected: "答二", Order: 2}, {ID: "p-1", CaseID: "case-1", ChapterID: "ch-1", Prompt: "第一问", Expected: "答一", Order: 1}}
	labels := PromptLabels(prompts)
	if len(labels) != 2 || labels[0] != "1. 第一问" {
		t.Fatalf("labels=%#v", labels)
	}
	points, max, err := ScoreRubric([]RubricCriterion{{ID: "logic", Label: "逻辑", Description: "论证", MaxPoints: 10}}, []RubricScore{{CriterionID: "logic", Points: 8}})
	if err != nil || points != 8 || max != 10 || RubricPercent(points, max) != 80 {
		t.Fatalf("score=%d max=%d err=%v", points, max, err)
	}
}

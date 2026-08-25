package domain

import (
	"fmt"
	"strings"
)

type LessonAudience string

const (
	AudienceStudent  LessonAudience = "student"
	AudienceTeacher  LessonAudience = "teacher"
	AudienceInternal LessonAudience = "internal"
)

type LessonPlan struct {
	ID           string         `json:"id"`
	CaseID       string         `json:"case_id"`
	LearningGoal string         `json:"learning_goal"`
	Audience     LessonAudience `json:"audience"`
	DurationMin  int            `json:"duration_min"`
	Activities   []string       `json:"activities"`
}

type CaseReference struct {
	ID       string `json:"id"`
	CaseID   string `json:"case_id"`
	Label    string `json:"label"`
	Source   string `json:"source"`
	Location string `json:"location"`
}

type ReviewRequest struct {
	ID          string `json:"id"`
	CaseID      string `json:"case_id"`
	Reviewer    string `json:"reviewer"`
	RequestedAt string `json:"requested_at"`
	Decision    string `json:"decision"`
	Feedback    string `json:"feedback"`
}

func (p LessonPlan) Validate() error {
	if strings.TrimSpace(p.ID) == "" || strings.TrimSpace(p.CaseID) == "" {
		return fmt.Errorf("lesson plan ids are required")
	}
	if strings.TrimSpace(p.LearningGoal) == "" {
		return fmt.Errorf("lesson learning goal is required")
	}
	if p.Audience != AudienceStudent && p.Audience != AudienceTeacher && p.Audience != AudienceInternal {
		return fmt.Errorf("unsupported lesson audience %q", p.Audience)
	}
	if p.DurationMin < 1 || p.DurationMin > 480 {
		return fmt.Errorf("lesson duration must be between 1 and 480 minutes")
	}
	if len(p.Activities) == 0 {
		return fmt.Errorf("lesson requires an activity")
	}
	for _, activity := range p.Activities {
		if strings.TrimSpace(activity) == "" {
			return fmt.Errorf("lesson activity cannot be empty")
		}
	}
	return nil
}

func (p LessonPlan) ActivityCount() int { return len(p.Activities) }

func (p LessonPlan) IsReadyForClass() bool {
	return p.Audience == AudienceStudent && p.DurationMin >= 15 && len(p.Activities) >= 2
}

func (r CaseReference) Validate() error {
	if strings.TrimSpace(r.ID) == "" || strings.TrimSpace(r.CaseID) == "" {
		return fmt.Errorf("reference ids are required")
	}
	if strings.TrimSpace(r.Label) == "" || strings.TrimSpace(r.Source) == "" {
		return fmt.Errorf("reference label and source are required")
	}
	return nil
}

func (r ReviewRequest) Validate() error {
	if strings.TrimSpace(r.ID) == "" || strings.TrimSpace(r.CaseID) == "" || strings.TrimSpace(r.Reviewer) == "" {
		return fmt.Errorf("review request identity is required")
	}
	if !ValidDate(r.RequestedAt) {
		return fmt.Errorf("review requested date must be YYYY-MM-DD")
	}
	if r.Decision != "" && r.Decision != "approved" && r.Decision != "changes_requested" {
		return fmt.Errorf("unsupported review decision %q", r.Decision)
	}
	return nil
}

func (r ReviewRequest) IsResolved() bool {
	return r.Decision == "approved" || r.Decision == "changes_requested"
}

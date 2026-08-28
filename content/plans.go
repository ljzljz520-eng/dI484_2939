package content

import (
	"fmt"
	"strings"

	"example.com/casescript/domain"
	"example.com/casescript/store"
)

type PlanManager struct{ repo *store.Repository }

func NewPlanManager(repo *store.Repository) *PlanManager { return &PlanManager{repo: repo} }

func (m *PlanManager) SavePlan(plan domain.LessonPlan) error {
	plan.LearningGoal = strings.TrimSpace(plan.LearningGoal)
	activities := make([]string, 0, len(plan.Activities))
	for _, activity := range plan.Activities {
		if cleaned := strings.TrimSpace(activity); cleaned != "" {
			activities = append(activities, cleaned)
		}
	}
	plan.Activities = activities
	if err := plan.Validate(); err != nil {
		return err
	}
	return m.repo.PutLessonPlan(plan)
}

func (m *PlanManager) Plan(caseID string) (domain.LessonPlan, error) {
	plans, err := m.repo.ListLessonPlans(caseID)
	if err != nil {
		return domain.LessonPlan{}, err
	}
	if len(plans) == 0 {
		return domain.LessonPlan{}, fmt.Errorf("lesson plan for %s not found", caseID)
	}
	return plans[0], nil
}

func (m *PlanManager) IsClassReady(caseID string) (bool, error) {
	plan, err := m.Plan(caseID)
	if err != nil {
		return false, err
	}
	return plan.IsReadyForClass(), nil
}

func (m *PlanManager) AddActivity(caseID, activity string) (domain.LessonPlan, error) {
	plan, err := m.Plan(caseID)
	if err != nil {
		return domain.LessonPlan{}, err
	}
	activity = strings.TrimSpace(activity)
	if activity == "" {
		return domain.LessonPlan{}, fmt.Errorf("activity is required")
	}
	plan.Activities = append(plan.Activities, activity)
	if err := m.SavePlan(plan); err != nil {
		return domain.LessonPlan{}, err
	}
	return plan, nil
}

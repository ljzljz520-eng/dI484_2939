package store

import "example.com/casescript/domain"

func (r *Repository) PutLessonPlan(value domain.LessonPlan) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(lessonBucket, value.ID, value)
}

func (r *Repository) GetLessonPlan(id string) (domain.LessonPlan, error) {
	var value domain.LessonPlan
	return value, r.get(lessonBucket, id, &value)
}

func (r *Repository) ListLessonPlans(caseID string) ([]domain.LessonPlan, error) {
	items := make([]domain.LessonPlan, 0)
	err := r.list(lessonBucket, func() any { return &domain.LessonPlan{} }, func(item any) {
		value := *(item.(*domain.LessonPlan))
		if caseID == "" || value.CaseID == caseID {
			items = append(items, value)
		}
	})
	return items, err
}

func (r *Repository) DeleteLessonPlan(id string) error {
	return r.delete(lessonBucket, id)
}

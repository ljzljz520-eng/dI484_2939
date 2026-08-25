package store

import "example.com/casescript/domain"

func (r *Repository) PutChapter(value domain.Chapter) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(chapterBucket, value.ID, value)
}

func (r *Repository) GetChapter(id string) (domain.Chapter, error) {
	var value domain.Chapter
	return value, r.get(chapterBucket, id, &value)
}

func (r *Repository) ListChapters(caseID string) ([]domain.Chapter, error) {
	items := make([]domain.Chapter, 0)
	err := r.list(chapterBucket, func() any { return &domain.Chapter{} }, func(item any) {
		value := *(item.(*domain.Chapter))
		if caseID == "" || value.CaseID == caseID {
			items = append(items, value)
		}
	})
	return items, err
}

func (r *Repository) PutNote(value domain.ClassroomNote) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(noteBucket, value.ID, value)
}

func (r *Repository) ListNotes(caseID string) ([]domain.ClassroomNote, error) {
	items := make([]domain.ClassroomNote, 0)
	err := r.list(noteBucket, func() any { return &domain.ClassroomNote{} }, func(item any) {
		value := *(item.(*domain.ClassroomNote))
		if caseID == "" || value.CaseID == caseID {
			items = append(items, value)
		}
	})
	return items, err
}

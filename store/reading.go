package store

import "example.com/casescript/domain"

func (r *Repository) PutReadingPacket(value domain.ReadingPacket) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(readingBucket, value.ID, value)
}

func (r *Repository) GetReadingPacket(id string) (domain.ReadingPacket, error) {
	var value domain.ReadingPacket
	return value, r.get(readingBucket, id, &value)
}

func (r *Repository) ListReadingPackets(caseID string) ([]domain.ReadingPacket, error) {
	items := make([]domain.ReadingPacket, 0)
	err := r.list(readingBucket, func() any { return &domain.ReadingPacket{} }, func(item any) {
		value := *(item.(*domain.ReadingPacket))
		if caseID == "" || value.CaseID == caseID {
			items = append(items, value)
		}
	})
	return items, err
}

func (r *Repository) PutVocabulary(value domain.VocabularyEntry) error {
	if err := value.Validate(); err != nil {
		return err
	}
	return r.put(vocabularyBucket, value.Term+":"+value.ChapterID, value)
}

func (r *Repository) ListVocabulary(chapterID string) ([]domain.VocabularyEntry, error) {
	items := make([]domain.VocabularyEntry, 0)
	err := r.list(vocabularyBucket, func() any { return &domain.VocabularyEntry{} }, func(item any) {
		value := *(item.(*domain.VocabularyEntry))
		if chapterID == "" || value.ChapterID == chapterID {
			items = append(items, value)
		}
	})
	return items, err
}

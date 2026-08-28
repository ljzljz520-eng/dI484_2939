package publish

import (
	"fmt"
	"sort"
	"strings"

	"example.com/casescript/domain"
)

type PublicationManifest struct {
	CaseID       string
	Version      int
	PublishedAt  string
	ChapterIDs   []string
	ReferenceIDs []string
	Checksum     string
}

func NewManifest(value domain.Publication, chapters []domain.Chapter, references []domain.CaseReference) PublicationManifest {
	chapterIDs := make([]string, 0, len(chapters))
	for _, chapter := range chapters {
		chapterIDs = append(chapterIDs, chapter.ID)
	}
	referenceIDs := make([]string, 0, len(references))
	for _, reference := range references {
		referenceIDs = append(referenceIDs, reference.ID)
	}
	sort.Strings(chapterIDs)
	sort.Strings(referenceIDs)
	checksum := strings.Join(append(append([]string{value.CaseID, fmt.Sprint(value.Version), value.PublishedAt}, chapterIDs...), referenceIDs...), "|")
	return PublicationManifest{CaseID: value.CaseID, Version: value.Version, PublishedAt: value.PublishedAt, ChapterIDs: chapterIDs, ReferenceIDs: referenceIDs, Checksum: checksum}
}

func (m PublicationManifest) Valid() bool {
	return m.CaseID != "" && m.Version > 0 && domain.ValidDate(m.PublishedAt) && m.Checksum != ""
}

func (m PublicationManifest) ItemCount() int { return len(m.ChapterIDs) + len(m.ReferenceIDs) }

func (m PublicationManifest) Description() string {
	return fmt.Sprintf("%s v%d (%d items)", m.CaseID, m.Version, m.ItemCount())
}

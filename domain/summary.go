package domain

import (
	"fmt"
	"strings"
)

type CaseSummary struct {
	CaseID           string
	ChapterCount     int
	NoteCount        int
	ReferenceCount   int
	IsPublished      bool
	HasLessonPlan    bool
	LastReviewedDate string
}

func (s CaseSummary) CompletionPercent() int {
	checks := 0
	if s.ChapterCount > 0 {
		checks++
	}
	if s.NoteCount > 0 {
		checks++
	}
	if s.ReferenceCount > 0 {
		checks++
	}
	if s.IsPublished {
		checks++
	}
	if s.HasLessonPlan {
		checks++
	}
	return checks * 20
}

func (s CaseSummary) Label() string {
	state := "草稿"
	if s.IsPublished {
		state = "已发布"
	}
	return fmt.Sprintf("%s · %s · %d%%", s.CaseID, state, s.CompletionPercent())
}

func CompactSummary(value string, limit int) string {
	value = strings.Join(strings.Fields(value), " ")
	if limit <= 0 || len(value) <= limit {
		return value
	}
	if limit < 4 {
		return value[:limit]
	}
	return value[:limit-3] + "..."
}

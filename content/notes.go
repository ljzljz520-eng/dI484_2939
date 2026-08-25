package content

import (
	"fmt"
	"strings"

	"example.com/casescript/domain"
)

func RenderNotes(notes []domain.ClassroomNote) string {
	if len(notes) == 0 {
		return ""
	}
	var builder strings.Builder
	for index, note := range notes {
		if index > 0 {
			builder.WriteString("\n")
		}
		builder.WriteString(fmt.Sprintf("[%s] %s: %s", note.UpdatedAt, note.Author, note.Text))
	}
	return builder.String()
}

func NotesForChapter(notes []domain.ClassroomNote, chapterID string) []domain.ClassroomNote {
	result := make([]domain.ClassroomNote, 0)
	for _, note := range notes {
		if note.ChapterID == chapterID {
			result = append(result, note)
		}
	}
	return result
}

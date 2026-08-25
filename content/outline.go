package content

import (
	"fmt"
	"strings"

	"example.com/casescript/domain"
)

func RenderOutline(chapters []domain.Chapter) string {
	var builder strings.Builder
	for index, chapter := range chapters {
		if index > 0 {
			builder.WriteString("\n")
		}
		builder.WriteString(fmt.Sprintf("%d. %s\n%s", chapter.Position, chapter.Title, chapter.Body))
	}
	return builder.String()
}

func ChapterTitles(chapters []domain.Chapter) []string {
	result := make([]string, 0, len(chapters))
	for _, chapter := range chapters {
		result = append(result, chapter.Title)
	}
	return result
}

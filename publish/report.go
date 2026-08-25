package publish

import (
	"fmt"
	"strings"

	"example.com/casescript/domain"
)

func PublicationSummary(value domain.Publication) string {
	return fmt.Sprintf("版本 %d · %s · %s", value.Version, domain.StatusLabel(value.Status), value.PublishedAt)
}

func HandoutSections(handout string) []string {
	parts := strings.Split(handout, "\n\n")
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		if strings.TrimSpace(part) != "" {
			result = append(result, part)
		}
	}
	return result
}

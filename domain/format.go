package domain

import (
	"fmt"
	"strings"
)

func NormalizeText(value string) string {
	return strings.Join(strings.Fields(value), " ")
}

func CaseLabel(c LegalCase) string {
	return fmt.Sprintf("%s · %s", c.Number, c.Title)
}

func StatusLabel(status CaseStatus) string {
	switch status {
	case StatusDraft:
		return "草稿"
	case StatusPublished:
		return "已发布"
	case StatusArchived:
		return "已归档"
	default:
		return "未知"
	}
}

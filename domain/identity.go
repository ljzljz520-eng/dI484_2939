package domain

import (
	"fmt"
	"strings"
)

func NormalizeIdentifier(value string) string {
	value = strings.TrimSpace(strings.ToLower(value))
	value = strings.ReplaceAll(value, " ", "-")
	value = strings.ReplaceAll(value, "_", "-")
	return value
}

func CaseNumberKey(number string) string {
	parts := strings.FieldsFunc(strings.TrimSpace(number), func(r rune) bool { return r == '-' || r == '/' || r == ' ' })
	if len(parts) == 0 {
		return ""
	}
	return strings.Join(parts, ":")
}

func EnsureCaseNumber(number string) error {
	normalized := CaseNumberKey(number)
	if normalized == "" {
		return fmt.Errorf("case number cannot be empty")
	}
	for _, part := range strings.Split(normalized, ":") {
		if part == "" {
			return fmt.Errorf("case number contains an empty segment")
		}
	}
	return nil
}

func SameDate(left, right LegalCase) bool {
	return left.PublishDate != "" && left.PublishDate == right.PublishDate
}

func SortCaseNumbers(items []LegalCase) []LegalCase {
	result := append([]LegalCase(nil), items...)
	for i := 1; i < len(result); i++ {
		current := result[i]
		j := i - 1
		for j >= 0 && CaseNumberKey(result[j].Number) > CaseNumberKey(current.Number) {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = current
	}
	return result
}

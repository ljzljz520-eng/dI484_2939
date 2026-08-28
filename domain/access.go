package domain

import "strings"

type AccessRole string

const (
	RoleTeacher  AccessRole = "teacher"
	RoleStudent  AccessRole = "student"
	RoleReviewer AccessRole = "reviewer"
)

func CanEdit(role AccessRole, value LegalCase) bool {
	if role != RoleTeacher {
		return false
	}
	return value.Status == StatusDraft
}

func CanPublish(role AccessRole, value LegalCase) bool {
	return role == RoleTeacher && value.Status == StatusDraft
}

func CanRead(role AccessRole, value LegalCase) bool {
	if role == RoleStudent {
		return value.Status == StatusPublished
	}
	if role == RoleTeacher || role == RoleReviewer {
		return value.Status != StatusArchived
	}
	return false
}

func RoleLabel(role AccessRole) string {
	switch role {
	case RoleTeacher:
		return "教师"
	case RoleStudent:
		return "学生"
	case RoleReviewer:
		return "审阅人"
	default:
		return strings.TrimSpace(string(role))
	}
}

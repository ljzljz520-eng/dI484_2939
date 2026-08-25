package domain

import (
	"fmt"
	"strings"
)

type StudentFeedback struct {
	ID       string `json:"id"`
	CaseID   string `json:"case_id"`
	Student  string `json:"student"`
	Text     string `json:"text"`
	Date     string `json:"date"`
	Resolved bool   `json:"resolved"`
}

func (f StudentFeedback) Validate() error {
	if f.ID == "" || f.CaseID == "" || strings.TrimSpace(f.Student) == "" || strings.TrimSpace(f.Text) == "" {
		return fmt.Errorf("feedback identity and text are required")
	}
	if !ValidDate(f.Date) {
		return fmt.Errorf("feedback date must be YYYY-MM-DD")
	}
	return nil
}

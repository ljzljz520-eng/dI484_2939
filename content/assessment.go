package content

import (
	"fmt"
	"sort"
	"strings"
)

type DiscussionPrompt struct {
	ID        string
	CaseID    string
	ChapterID string
	Prompt    string
	Expected  string
	Order     int
}

type RubricCriterion struct {
	ID          string
	Label       string
	Description string
	MaxPoints   int
}

type RubricScore struct {
	CriterionID string
	Points      int
	Feedback    string
}

func (p DiscussionPrompt) Validate() error {
	if p.ID == "" || p.CaseID == "" || p.ChapterID == "" {
		return fmt.Errorf("discussion prompt ids are required")
	}
	if strings.TrimSpace(p.Prompt) == "" || strings.TrimSpace(p.Expected) == "" {
		return fmt.Errorf("discussion prompt text is required")
	}
	if p.Order < 1 {
		return fmt.Errorf("discussion prompt order must be positive")
	}
	return nil
}

func SortPrompts(prompts []DiscussionPrompt) []DiscussionPrompt {
	result := append([]DiscussionPrompt(nil), prompts...)
	sort.SliceStable(result, func(i, j int) bool {
		if result[i].Order != result[j].Order {
			return result[i].Order < result[j].Order
		}
		return result[i].ID < result[j].ID
	})
	return result
}

func PromptLabels(prompts []DiscussionPrompt) []string {
	ordered := SortPrompts(prompts)
	result := make([]string, 0, len(ordered))
	for _, prompt := range ordered {
		result = append(result, fmt.Sprintf("%d. %s", prompt.Order, prompt.Prompt))
	}
	return result
}

func (c RubricCriterion) Validate() error {
	if c.ID == "" || strings.TrimSpace(c.Label) == "" || strings.TrimSpace(c.Description) == "" {
		return fmt.Errorf("rubric criterion is incomplete")
	}
	if c.MaxPoints < 1 || c.MaxPoints > 100 {
		return fmt.Errorf("rubric points must be between 1 and 100")
	}
	return nil
}

func ScoreRubric(criteria []RubricCriterion, scores []RubricScore) (int, int, error) {
	max := 0
	points := 0
	limits := make(map[string]int)
	for _, criterion := range criteria {
		if err := criterion.Validate(); err != nil {
			return 0, 0, err
		}
		max += criterion.MaxPoints
		limits[criterion.ID] = criterion.MaxPoints
	}
	for _, score := range scores {
		limit, ok := limits[score.CriterionID]
		if !ok {
			return 0, 0, fmt.Errorf("unknown rubric criterion %s", score.CriterionID)
		}
		if score.Points < 0 || score.Points > limit {
			return 0, 0, fmt.Errorf("score for %s is out of range", score.CriterionID)
		}
		points += score.Points
	}
	return points, max, nil
}

func RubricPercent(points, max int) int {
	if max <= 0 || points <= 0 {
		return 0
	}
	return points * 100 / max
}

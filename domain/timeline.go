package domain

import "sort"

type TimelineEvent struct {
	ID       string `json:"id"`
	CaseID   string `json:"case_id"`
	Date     string `json:"date"`
	Kind     string `json:"kind"`
	Actor    string `json:"actor"`
	Message  string `json:"message"`
	Sequence int    `json:"sequence"`
}

func (e TimelineEvent) Validate() error {
	if e.ID == "" || e.CaseID == "" || e.Kind == "" || e.Actor == "" || e.Message == "" {
		return errRequiredTimeline
	}
	if !ValidDate(e.Date) || e.Sequence < 1 {
		return errInvalidTimeline
	}
	return nil
}

var errRequiredTimeline = timelineError("timeline event fields are required")
var errInvalidTimeline = timelineError("timeline event date and sequence are required")

type timelineError string

func (e timelineError) Error() string { return string(e) }

func SortTimeline(events []TimelineEvent) []TimelineEvent {
	result := append([]TimelineEvent(nil), events...)
	sort.SliceStable(result, func(i, j int) bool {
		if result[i].Date != result[j].Date {
			return result[i].Date < result[j].Date
		}
		if result[i].Sequence != result[j].Sequence {
			return result[i].Sequence < result[j].Sequence
		}
		return result[i].ID < result[j].ID
	})
	return result
}

func TimelineKinds(events []TimelineEvent) map[string]int {
	counts := make(map[string]int)
	for _, event := range events {
		counts[event.Kind]++
	}
	return counts
}

package domain

import "sort"

type ProgressMark struct {
	CaseID     string
	Student    string
	Completed  []string
	LastOpened string
	Percent    int
}

func (p ProgressMark) Recalculate(total int) ProgressMark {
	if total <= 0 {
		p.Percent = 0
		return p
	}
	p.Percent = len(p.Completed) * 100 / total
	if p.Percent > 100 {
		p.Percent = 100
	}
	return p
}

func (p ProgressMark) Finished(sectionID string) bool {
	for _, completed := range p.Completed {
		if completed == sectionID {
			return true
		}
	}
	return false
}

func (p *ProgressMark) Mark(sectionID string) {
	if sectionID == "" || p.Finished(sectionID) {
		return
	}
	p.Completed = append(p.Completed, sectionID)
	sort.Strings(p.Completed)
}

func CompletionLabel(percent int) string {
	switch {
	case percent >= 100:
		return "已完成"
	case percent >= 50:
		return "进行中"
	default:
		return "未开始"
	}
}

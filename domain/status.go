package domain

type Status string

func StatusPublished() Status {
	return Status("published")
}

func StatusDraft() Status {
	return Status("draft")
}

func NewStatus(s string) Status {
	switch s {
	case "published":
		return StatusPublished()
	case "draft":
		return StatusDraft()
	default:
		return StatusDraft()
	}
}

func (status Status) string() string {
	return string(status)
}

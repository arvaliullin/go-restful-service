package domain

type GlossaryTerm struct {
	ID          int64  `json:"id"`
	Term        string `json:"term"`
	Description string `json:"description"`
}

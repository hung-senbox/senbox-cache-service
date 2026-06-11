package term

type TermCache struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	Color         string     `json:"color"`
	IsCurrentTerm bool       `json:"is_current_term"`
	StartDate     string     `json:"start_date"`
	EndDate       string     `json:"end_date"`
	TermWeeks     []TermWeek `json:"term_weeks"`
}

type TermWeek struct {
	WeekNumber    int    `json:"week_number"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	ResourceCount int    `json:"resource_count"`
}

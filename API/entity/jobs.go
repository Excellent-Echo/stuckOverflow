package entity

type Data struct {
	Jobs []Jobs `json:"jobs"`
}

type Jobs struct {
	ID                        uint32 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Title                     string `json:"title"`
	Url                       string `json:"url"`
	CompanyName               string `json:"company_name"`
	JobType                   string `json:"job_type"`
	PublicationDate           string `json:"publication_date"`
	Description               string `json:"description"`
	CandidateRequiredLocation string `json:"candidate_required_location"`
}

type Pagination struct {
	Limit  int    `json:"limit"`
	Page   int    `json:"page"`
	Sort   string `json:"sort"`
	Search string `json:"search"`
}

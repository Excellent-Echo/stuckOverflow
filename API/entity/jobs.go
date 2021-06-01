package entity

type Data struct {
	Jobs []Jobs `json:"jobs"`
}

type Jobs struct {
	ID              uint32 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	Title           string `json:"title"`
	Url             string `json:"url"`
	CompanyName     string `json:"company_name"`
	JobType         string `json:"job_type"`
	PublicationDate string `json:"publication_date"`
}

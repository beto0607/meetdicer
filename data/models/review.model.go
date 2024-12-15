package models

type ReviewModel struct {
	Id             string
	UserId         string
	Content        string
	ContentHistory []string
	Date           string
	CreatedAt      string
	DeletedAt      string
	ModifiedAt     string
}

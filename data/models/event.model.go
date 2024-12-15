package models

type EventModel struct {
	Id         string
	VenueId    string
	GameIds    []string
	Date       string
	Comments   []CommentModel
	CreatedAt  string
	DeletedAt  string
	ModifiedAt string
}

type CommentModel struct {
	Id             string
	Content        string
	ContentHistory []string
	UserId         string
	Private        bool
	Reactions      []string
	CreatedAt      string
	DeletedAt      string
	ModifiedAt     string
}

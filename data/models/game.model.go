package models

type GameModel struct {
	Id                 string
	Name               string
	MinNumberOfPlayers int
	MaxNumberOfPlayers int
	Description        string
	TimesPlayed        int
	Links              struct {
		PublisherLink string
		BGG           string
		Rules         string
	}
	CreatedAt  string
	DeletedAt  string
	ModifiedAt string
}

package models

type VenueModel struct {
	Id          string
	Name        string
	NameHistory []string
	Public      bool
	Location    struct {
		Latitude  string
		Longiture string
		Address   string
		City      string
		Country   string
	}
	Links struct {
		Website    string
		Email      string
		Instagram  string
		Facebook   string
		GoogleMaps string
	}
	GamesPlayed []string
	CreatedAt   string
	DeletedAt   string
	ModifiedAt  string
}

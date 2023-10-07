package domain

type Account struct {
	ID           string
	Name         string
	LastName     string
	Email        string
	Mobile       string
	Language     string
	Country      string
	TypeAccount  string
	Password     string
	TypeStatus   string
	CreationDate string
	Sub          *string
}

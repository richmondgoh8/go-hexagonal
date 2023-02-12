package domain

type Link struct {
	ID   string `json:"id" db:"id"`
	Url  string `json:"url" db:"url"`
	Name string `json:"name" db:"name"`
}

package models

type Recital struct {
	ID     string   `json:"id"`
	Number int      `json:"number"`
	Texts  []string `json:"texts"`
}

package models

type Article struct {
	ID                 string `json:"id"`
	Number             int    `json:"number"`
	Roman              string `json:"roman"`
	Title              string `json:"title"`
	NumberOfParagraphs int    `json:"number_of_paragraphs"`
}

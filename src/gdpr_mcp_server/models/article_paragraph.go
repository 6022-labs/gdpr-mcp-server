package models

type ArticleParagraph struct {
	Number    int      `json:"number"`
	ArticleId string   `json:"article_id"`
	Texts     []string `json:"texts"`
}

package models

type Chapter struct {
	ID          string   `json:"id"`
	Roman       string   `json:"roman"`
	Number      int      `json:"number"`
	Title       string   `json:"title"`
	ArticlesIds []string `json:"articles_ids"`
}

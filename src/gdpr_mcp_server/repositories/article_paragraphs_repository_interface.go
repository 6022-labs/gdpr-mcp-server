package repositories

import "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"

type ArticleParagraphsRepositoryInterface interface {
	GetByArticleIdAndIndex(articleId string, index uint) (*models.ArticleParagraph, error)
}

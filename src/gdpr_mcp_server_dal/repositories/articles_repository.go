package repositories

import (
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"
)

type ArticlesRepository struct {
	gdprDataClient gdpr_mcp_server_dal.GdprDataClientInterface
}

func NewArticlesRepository(
	gdprDataClient gdpr_mcp_server_dal.GdprDataClientInterface,
) *ArticlesRepository {
	return &ArticlesRepository{
		gdprDataClient: gdprDataClient,
	}
}

func (r *ArticlesRepository) GetById(articleId string) (*models.Article, error) {
	articleSet := r.gdprDataClient.ArticlesSetSnapshot()
	if article, exists := articleSet[articleId]; exists {
		return article, nil
	}

	return nil, nil
}

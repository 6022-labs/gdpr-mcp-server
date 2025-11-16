package repositories

import (
	"fmt"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"
)

type ArticleParagraphsRepository struct {
	gdprDataClient gdpr_mcp_server_dal.GdprDataClientInterface
}

func NewArticleParagraphsRepository(
	gdprDataClient gdpr_mcp_server_dal.GdprDataClientInterface,
) *ArticleParagraphsRepository {
	return &ArticleParagraphsRepository{
		gdprDataClient: gdprDataClient,
	}
}

func (r *ArticleParagraphsRepository) GetByArticleIdAndIndex(articleId string, index uint) (*models.ArticleParagraph, error) {
	articleParagraphSet := r.gdprDataClient.ArticleParagraphsSetSnapshot()
	if articleParagraphs, exists := articleParagraphSet[articleId]; exists {
		if int(index) < len(articleParagraphs) {
			return articleParagraphs[index], nil
		}

		return nil, fmt.Errorf("index out of range")
	}

	return nil, nil
}

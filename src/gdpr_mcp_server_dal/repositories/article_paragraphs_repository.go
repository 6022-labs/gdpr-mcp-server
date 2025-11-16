package repositories

import "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"

type ArticleParagraphsRepository struct {
	gdprDataClient *gdpr_mcp_server_dal.GdprDataClient
}

func NewArticleParagraphsRepository(
	gdprDataClient *gdpr_mcp_server_dal.GdprDataClient,
) *ArticleParagraphsRepository {
	return &ArticleParagraphsRepository{
		gdprDataClient: gdprDataClient,
	}
}

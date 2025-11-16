package repositories

import "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"

type ArticlesRepository struct {
	gdprDataClient *gdpr_mcp_server_dal.GdprDataClient
}

func NewArticlesRepository(
	gdprDataClient *gdpr_mcp_server_dal.GdprDataClient,
) *ArticlesRepository {
	return &ArticlesRepository{
		gdprDataClient: gdprDataClient,
	}
}

package repositories

import "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"

type ChaptersRepository struct {
	gdprDataClient *gdpr_mcp_server_dal.GdprDataClient
}

func NewChaptersRepository(
	gdprDataClient *gdpr_mcp_server_dal.GdprDataClient,
) *ChaptersRepository {
	return &ChaptersRepository{
		gdprDataClient: gdprDataClient,
	}
}

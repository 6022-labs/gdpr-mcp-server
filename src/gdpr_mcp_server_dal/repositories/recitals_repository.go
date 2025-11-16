package repositories

import "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"

type RecitalsRepository struct {
	gdprDataClient *gdpr_mcp_server_dal.GdprDataClient
}

func NewRecitalsRepository(
	gdprDataClient *gdpr_mcp_server_dal.GdprDataClient,
) *RecitalsRepository {
	return &RecitalsRepository{
		gdprDataClient: gdprDataClient,
	}
}

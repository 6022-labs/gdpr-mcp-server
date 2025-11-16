package repositories

import (
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"
)

type RecitalsRepository struct {
	gdprDataClient gdpr_mcp_server_dal.GdprDataClientInterface
}

func NewRecitalsRepository(
	gdprDataClient gdpr_mcp_server_dal.GdprDataClientInterface,
) *RecitalsRepository {
	return &RecitalsRepository{
		gdprDataClient: gdprDataClient,
	}
}

func (r *RecitalsRepository) GetById(recitalId string) (*models.Recital, error) {
	recitalSet := r.gdprDataClient.RecitalsSetSnapshot()
	if recital, exists := recitalSet[recitalId]; exists {
		return recital, nil
	}

	return nil, nil
}

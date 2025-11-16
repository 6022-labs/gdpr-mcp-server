package repositories

import "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"

type RecitalsRepositoryInterface interface {
	GetById(recitalId string) (*models.Recital, error)
}

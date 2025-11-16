package repositories

import "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"

type ChaptersRepositoryInterface interface {
	GetById(chapterId string) (*models.Chapter, error)
}

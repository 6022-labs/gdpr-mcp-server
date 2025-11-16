package repositories

import (
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"
)

type ChaptersRepository struct {
	gdprDataClient gdpr_mcp_server_dal.GdprDataClientInterface
}

func NewChaptersRepository(
	gdprDataClient gdpr_mcp_server_dal.GdprDataClientInterface,
) *ChaptersRepository {
	return &ChaptersRepository{
		gdprDataClient: gdprDataClient,
	}
}

func (r *ChaptersRepository) GetById(chapterId string) (*models.Chapter, error) {
	chapterSet := r.gdprDataClient.ChaptersSetSnapshot()
	if chapter, exists := chapterSet[chapterId]; exists {
		return chapter, nil
	}

	return nil, nil
}

package gdpr_mcp_server_dal

import (
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
)

type GdprDataClientInterface interface {
	RecitalsSetSnapshot() map[string]*models.Recital
	ChaptersSetSnapshot() map[string]*models.Chapter
	ArticlesSetSnapshot() map[string]*models.Article
	ArticleParagraphsSetSnapshot() map[string][]*models.ArticleParagraph
}

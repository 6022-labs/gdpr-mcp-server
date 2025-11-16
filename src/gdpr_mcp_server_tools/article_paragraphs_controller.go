package gdpr_mcp_server_tools

import (
	"context"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/repositories"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

type ArticleParagraphsController struct {
	logger                      *zap.Logger
	articleParagraphsRepository repositories.ArticleParagraphsRepositoryInterface
}

func NewArticleParagraphsController(
	logger *zap.Logger,
	articleParagraphsRepository repositories.ArticleParagraphsRepositoryInterface,
) *ArticleParagraphsController {
	return &ArticleParagraphsController{
		logger:                      logger,
		articleParagraphsRepository: articleParagraphsRepository,
	}
}

func (c *ArticleParagraphsController) RegisterTools(mcpServer *mcp.Server) {
	mcp.AddTool(mcpServer, &mcp.Tool{Name: "GetArticleParagraphsByArticleId", Description: "Get a paragraph for a given article ID (art-1, art-2, ...) and index"}, c.GetArticleParagraphsByArticleId)
}

type GetArticleParagraphsByArticleIdInput struct {
	ArticleId string `json:"article_id"`
	Index     uint   `json:"index"`
}

func (c *ArticleParagraphsController) GetArticleParagraphsByArticleId(ctx context.Context, req *mcp.CallToolRequest, input GetArticleParagraphsByArticleIdInput) (
	*mcp.CallToolResult,
	*models.ArticleParagraph,
	error,
) {
	paragraph, err := c.articleParagraphsRepository.GetByArticleIdAndIndex(input.ArticleId, input.Index)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{}, paragraph, nil
}

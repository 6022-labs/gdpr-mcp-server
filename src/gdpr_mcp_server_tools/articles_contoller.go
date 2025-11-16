package gdpr_mcp_server_tools

import (
	"context"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/repositories"
	"go.uber.org/zap"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type ArticlesController struct {
	logger              *zap.Logger
	articleRepositories repositories.ArticlesRepositoryInterface
}

func NewArticlesController(
	logger *zap.Logger,
	articleRepositories repositories.ArticlesRepositoryInterface,
) *ArticlesController {
	return &ArticlesController{
		logger:              logger,
		articleRepositories: articleRepositories,
	}
}

func (c *ArticlesController) RegisterTools(mcpServer *mcp.Server) {
	mcp.AddTool(mcpServer, &mcp.Tool{Name: "GetArticleById", Description: "Get a single GDPR article using its ID (art-1, art-2, etc...)"}, c.GetArticleById)
}

type GetArticleByIdInput struct {
	ArticleId string `json:"article_id"`
}

func (c *ArticlesController) GetArticleById(ctx context.Context, req *mcp.CallToolRequest, input GetArticleByIdInput) (
	*mcp.CallToolResult,
	*models.Article,
	error,
) {
	article, err := c.articleRepositories.GetById(input.ArticleId)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{}, article, nil
}

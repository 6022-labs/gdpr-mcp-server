package gdpr_mcp_server_tools

import (
	"context"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/repositories"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

type ChaptersController struct {
	logger              *zap.Logger
	chapterRepositories repositories.ChaptersRepositoryInterface
}

func NewChaptersController(
	logger *zap.Logger,
	chapterRepositories repositories.ChaptersRepositoryInterface,
) *ChaptersController {
	return &ChaptersController{
		logger:              logger,
		chapterRepositories: chapterRepositories,
	}
}

func (c *ChaptersController) RegisterTools(mcpServer *mcp.Server) {
	mcp.AddTool(mcpServer, &mcp.Tool{Name: "GetChapterById", Description: "Get a single GDPR chapter using its ID (chap-1, chap-2, etc...)"}, c.GetChapterById)
}

type GetChapterByIdInput struct {
	ChapterId string `json:"chapter_id"`
}

func (c *ChaptersController) GetChapterById(ctx context.Context, req *mcp.CallToolRequest, input GetChapterByIdInput) (
	*mcp.CallToolResult,
	*models.Chapter,
	error,
) {
	chapter, err := c.chapterRepositories.GetById(input.ChapterId)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{}, chapter, nil
}

package gdpr_mcp_server_tools

import (
	"context"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/repositories"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

type RecitalsController struct {
	logger              *zap.Logger
	recitalRepositories repositories.RecitalsRepositoryInterface
}

func NewRecitalsController(
	logger *zap.Logger,
	recitalRepositories repositories.RecitalsRepositoryInterface,
) *RecitalsController {
	return &RecitalsController{
		logger:              logger,
		recitalRepositories: recitalRepositories,
	}
}

func (c *RecitalsController) RegisterTools(mcpServer *mcp.Server) {
	mcp.AddTool(mcpServer, &mcp.Tool{Name: "GetRecitalById", Description: "Get a single GDPR recital using its ID (rec-1, rec-2, etc...)"}, c.GetRecitalById)
}

type GetRecitalByIdInput struct {
	RecitalId string `json:"recital_id"`
}

func (c *RecitalsController) GetRecitalById(ctx context.Context, req *mcp.CallToolRequest, input GetRecitalByIdInput) (
	*mcp.CallToolResult,
	*models.Recital,
	error,
) {
	recital, err := c.recitalRepositories.GetById(input.RecitalId)
	if err != nil {
		return nil, nil, err
	}

	return &mcp.CallToolResult{}, recital, nil
}

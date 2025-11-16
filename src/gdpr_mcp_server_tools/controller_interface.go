package gdpr_mcp_server_tools

import "github.com/modelcontextprotocol/go-sdk/mcp"

type ControllerInterface interface {
	RegisterTools(mcpServer *mcp.Server)
}

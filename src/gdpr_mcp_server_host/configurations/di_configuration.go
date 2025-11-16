package configurations

import (
	gdpr_mcp_server_configurations "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/configurations"
	gdpr_mcp_server_dal_configurations "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal/configurations"
	gdpr_mcp_server_host_middlewares "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_host/middlewares"
	gdpr_mcp_server_host_settings "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_host/settings"
	gdpr_mcp_server_tools_configurations "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_tools/configurations"
	"go.uber.org/dig"
)

func ConfigureDI() *dig.Container {
	container := dig.New()

	container.Provide(gdpr_mcp_server_host_settings.NewHostSettings)
	container.Provide(gdpr_mcp_server_host_middlewares.NewLoggingMiddleware)

	gdpr_mcp_server_configurations.AddGdprMcpServerConfiguration(container)
	gdpr_mcp_server_dal_configurations.AddGdprMcpServerDalConfiguration(container)
	gdpr_mcp_server_tools_configurations.AddGdprMcpServerToolsConfiguration(container)

	return container
}

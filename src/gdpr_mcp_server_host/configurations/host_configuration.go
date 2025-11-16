package configurations

import (
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_host/middlewares"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_host/settings"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func ConfigureHost(container *dig.Container) {
	container.Provide(newHttpMcpServer)

	err := container.Invoke(useTools)
	if err != nil {
		panic(err)
	}
}

func newHttpMcpServer(hostSettings *settings.HostSettings) *mcp.Server {
	return mcp.NewServer(&mcp.Implementation{Name: hostSettings.AppName, Version: "v1.0.0"}, nil)
}

type useToolsParams struct {
	dig.In

	Server            *mcp.Server
	Logger            *zap.Logger
	LoggingMiddleware *middlewares.LoggingMiddleware
	Controllers       []gdpr_mcp_server_tools.ControllerInterface `group:"controllers"`
}

func useTools(p useToolsParams) {
	p.Server.AddReceivingMiddleware(p.LoggingMiddleware.Handle)

	for _, controller := range p.Controllers {
		controller.RegisterTools(p.Server)
	}
}

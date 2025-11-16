package configurations

import (
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_tools"
	"go.uber.org/dig"
)

func AddGdprMcpServerToolsConfiguration(container *dig.Container) {
	// Controllers
	err := container.Provide(
		gdpr_mcp_server_tools.NewArticlesController,
		dig.As(new(gdpr_mcp_server_tools.ControllerInterface)),
		dig.Group("controllers"),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		gdpr_mcp_server_tools.NewChaptersController,
		dig.As(new(gdpr_mcp_server_tools.ControllerInterface)),
		dig.Group("controllers"),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		gdpr_mcp_server_tools.NewRecitalsController,
		dig.As(new(gdpr_mcp_server_tools.ControllerInterface)),
		dig.Group("controllers"),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		gdpr_mcp_server_tools.NewArticleParagraphsController,
		dig.As(new(gdpr_mcp_server_tools.ControllerInterface)),
		dig.Group("controllers"),
	)
	if err != nil {
		panic(err)
	}
}

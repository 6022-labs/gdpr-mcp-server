package configurations

import (
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/repositories"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"
	infra_repositories "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal/repositories"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal/settings"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

func AddGdprMcpServerDalConfiguration(container *dig.Container) {
	// Clients
	err := container.Provide(
		func(dataSettings *settings.DataSettings, logger *zap.Logger) *gdpr_mcp_server_dal.GdprDataClient {
			client, err := gdpr_mcp_server_dal.NewGdprDataClient(dataSettings, logger)
			if err != nil {
				panic(err)
			}

			return client
		},
	)
	if err != nil {
		panic(err)
	}

	// Settings
	err = container.Provide(
		settings.NewDataSettings,
	)
	if err != nil {
		panic(err)
	}

	// Repositories
	err = container.Provide(
		infra_repositories.NewArticlesRepository,
		dig.As(new(repositories.ArticlesRepositoryInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		infra_repositories.NewChaptersRepository,
		dig.As(new(repositories.ChaptersRepositoryInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		infra_repositories.NewRecitalsRepository,
		dig.As(new(repositories.RecitalsRepositoryInterface)),
	)
	if err != nil {
		panic(err)
	}

	err = container.Provide(
		infra_repositories.NewArticleParagraphsRepository,
		dig.As(new(repositories.ArticleParagraphsRepositoryInterface)),
	)
	if err != nil {
		panic(err)
	}
}

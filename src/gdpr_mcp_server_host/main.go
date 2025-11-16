package main

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_host/configurations"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_host/settings"
	"github.com/joho/godotenv"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	slogzap "github.com/samber/slog-zap"
	"go.uber.org/zap"
)

func main() {
	godotenv.Load(".env")

	container := configurations.ConfigureDI()
	configurations.ConfigureLogging(container)
	configurations.ConfigureHost(container)

	err := container.Invoke(func(server *mcp.Server, logger *zap.Logger, hostSettings *settings.HostSettings) error {
		slogLogger := slog.New(slogzap.Option{Level: slog.LevelDebug, Logger: logger}.NewZapHandler())

		handler := mcp.NewStreamableHTTPHandler(func(req *http.Request) *mcp.Server {
			return server
		}, &mcp.StreamableHTTPOptions{
			Stateless:    hostSettings.Stateless,
			JSONResponse: hostSettings.JSONResponse,
			Logger:       slogLogger,
		})

		logger.Info("Starting MCP server listening", zap.Int("port", hostSettings.ApiPort))

		return http.ListenAndServe(fmt.Sprintf(":%d", hostSettings.ApiPort), handler)
	})

	if err != nil {
		panic(err)
	}
}

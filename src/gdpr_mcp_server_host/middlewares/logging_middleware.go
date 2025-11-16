package middlewares

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"go.uber.org/zap"
)

type LoggingMiddleware struct {
	logger *zap.Logger
}

func NewLoggingMiddleware(logger *zap.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

func (lm *LoggingMiddleware) Handle(next mcp.MethodHandler) mcp.MethodHandler {
	return func(ctx context.Context, method string, req mcp.Request) (mcp.Result, error) {
		start := time.Now()
		sessionID := req.GetSession().ID()

		reqId := uuid.New().String()

		fields := []zap.Field{
			zap.String("event", "mcp_call_started"),
			zap.String("session_id", sessionID),
			zap.String("method", method),
			zap.String("request_id", reqId),
			zap.Time("ts", start),
		}
		lm.logger.Info("mcp request started", fields...)

		result, err := next(ctx, method, req)

		duration := time.Since(start)
		fields = []zap.Field{
			zap.String("session_id", sessionID),
			zap.String("method", method),
			zap.String("request_id", reqId),
			zap.Duration("duration", duration),
		}

		if err != nil {
			fields = append(fields,
				zap.Error(err),
				zap.String("event", "mcp_call_failed"),
			)
			lm.logger.Error("mcp request failed", fields...)
		} else {
			fields = append(fields,
				zap.String("event", "mcp_call_succeeded"),
			)
			lm.logger.Info("mcp request succeeded", fields...)
		}

		return result, err
	}
}

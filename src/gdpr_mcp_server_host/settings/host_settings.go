package settings

import (
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

type HostSettings struct {
	ApiPort      int
	AppName      string
	Stateless    bool
	JSONResponse bool
}

func NewHostSettings(logger *zap.Logger) *HostSettings {
	apiPortStr := os.Getenv("API_PORT")
	if len(strings.TrimSpace(apiPortStr)) == 0 {
		logger.Warn("environment variable API_PORT is not set, defaulting to 3000")
		apiPortStr = "3000"
	}

	apiPort, err := strconv.Atoi(apiPortStr)
	if err != nil {
		logger.Warn("invalid API_PORT environment variable value, defaulting to 3000")
		apiPort = 3000
	}

	appName := os.Getenv("APP_NAME")
	if len(strings.TrimSpace(appName)) == 0 {
		logger.Fatal("please set your APP_NAME value in your environment")
	}

	statelessStr := os.Getenv("STATELESS")
	stateless := false
	if len(strings.TrimSpace(statelessStr)) > 0 {
		stateless, _ = strconv.ParseBool(statelessStr)
	}

	jsonResponseStr := os.Getenv("JSON_RESPONSE")
	jsonResponse := false
	if len(strings.TrimSpace(jsonResponseStr)) > 0 {
		jsonResponse, _ = strconv.ParseBool(jsonResponseStr)
	}

	return &HostSettings{
		ApiPort:      apiPort,
		AppName:      appName,
		Stateless:    stateless,
		JSONResponse: jsonResponse,
	}
}

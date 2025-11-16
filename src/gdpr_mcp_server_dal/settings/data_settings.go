package settings

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

type DataSettings struct {
	ArticlesDataFilePath string
	ChaptersDataFilePath string
	RecitalsDataFilePath string
}

func NewDataSettings(logger *zap.Logger) *DataSettings {
	articlesDataFilePath := os.Getenv("DAL_ARTICLES_DATA_FILE_PATH")
	if len(strings.TrimSpace(articlesDataFilePath)) == 0 {
		logger.Fatal("please set your DAL_ARTICLES_DATA_FILE_PATH value in your environment")
	}

	chaptersDataFilePath := os.Getenv("DAL_CHAPTERS_DATA_FILE_PATH")
	if len(strings.TrimSpace(chaptersDataFilePath)) == 0 {
		logger.Fatal("please set your DAL_CHAPTERS_DATA_FILE_PATH value in your environment")
	}

	recitalsDataFilePath := os.Getenv("DAL_RECITALS_DATA_FILE_PATH")
	if len(strings.TrimSpace(recitalsDataFilePath)) == 0 {
		logger.Fatal("please set your DAL_RECITALS_DATA_FILE_PATH value in your environment")
	}

	return &DataSettings{
		ArticlesDataFilePath: articlesDataFilePath,
		ChaptersDataFilePath: chaptersDataFilePath,
		RecitalsDataFilePath: recitalsDataFilePath,
	}
}

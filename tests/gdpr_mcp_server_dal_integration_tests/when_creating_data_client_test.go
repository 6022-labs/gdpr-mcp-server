package gdpr_mcp_server_dal_integration_tests

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	dal "github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal/settings"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type WhenCreatingDataClientTestingSuite struct{}

func WhenCreatingDataClientBeforeEach() *WhenCreatingDataClientTestingSuite {
	return &WhenCreatingDataClientTestingSuite{}
}

func (s *WhenCreatingDataClientTestingSuite) repoRoot(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("failed to resolve caller path")
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
	abs, err := filepath.Abs(root)
	if err != nil {
		t.Fatalf("failed to compute absolute root: %v", err)
	}
	return abs
}

func (s *WhenCreatingDataClientTestingSuite) realDataSettings(t *testing.T) *settings.DataSettings {
	t.Helper()
	root := s.repoRoot(t)
	return &settings.DataSettings{
		ArticlesDataFilePath: filepath.Join(root, "data", "v1", "articles"),
		ChaptersDataFilePath: filepath.Join(root, "data", "v1", "chapters"),
		RecitalsDataFilePath: filepath.Join(root, "data", "v1", "recitals"),
	}
}

func (s *WhenCreatingDataClientTestingSuite) emptyTempDataSettings(t *testing.T) *settings.DataSettings {
	t.Helper()
	base := t.TempDir()
	arts := filepath.Join(base, "articles")
	chs := filepath.Join(base, "chapters")
	recs := filepath.Join(base, "recitals")
	for _, d := range []string{arts, chs, recs} {
		assert.NoError(t, os.MkdirAll(d, 0o755))
	}
	return &settings.DataSettings{
		ArticlesDataFilePath: arts,
		ChaptersDataFilePath: chs,
		RecitalsDataFilePath: recs,
	}
}

func (s *WhenCreatingDataClientTestingSuite) badRecitalTempDataSettings(t *testing.T) *settings.DataSettings {
	t.Helper()
	base := t.TempDir()
	arts := filepath.Join(base, "articles")
	chs := filepath.Join(base, "chapters")
	recs := filepath.Join(base, "recitals")
	for _, d := range []string{arts, chs, recs} {
		assert.NoError(t, os.MkdirAll(d, 0o755))
	}

	// Create a recital JSON missing the required ID to trigger an error path.
	f := filepath.Join(recs, "rec-bad.json")
	assert.NoError(t, os.WriteFile(f, []byte(`{"number":1,"texts":["x"]}`), 0o644))
	return &settings.DataSettings{
		ArticlesDataFilePath: arts,
		ChaptersDataFilePath: chs,
		RecitalsDataFilePath: recs,
	}
}

func TestWhenCreatingDataClient(t *testing.T) {
	suite := WhenCreatingDataClientBeforeEach()

	t.Run("Given real GDPR dataset on disk", func(t *testing.T) {
		ds := suite.realDataSettings(t)
		logger := zap.NewNop()

		t.Run("Should construct client and load non-empty sets", func(t *testing.T) {
			cli, err := dal.NewGdprDataClient(ds, logger)

			assert.NoError(t, err)
			assert.NotNil(t, cli)

			recs := cli.RecitalsSetSnapshot()
			chs := cli.ChaptersSetSnapshot()
			arts := cli.ArticlesSetSnapshot()
			paras := cli.ArticleParagraphsSetSnapshot()

			assert.Greater(t, len(recs), 0, "recitals should not be empty")
			assert.Greater(t, len(chs), 0, "chapters should not be empty")
			assert.Greater(t, len(arts), 0, "articles should not be empty")
			assert.Greater(t, len(paras), 0, "article paragraphs should not be empty")
		})
	})

	t.Run("Given empty data directories", func(t *testing.T) {
		ds := suite.emptyTempDataSettings(t)
		logger := zap.NewNop()

		t.Run("Should construct client with zero counts and no error", func(t *testing.T) {
			cli, err := dal.NewGdprDataClient(ds, logger)
			assert.NoError(t, err)
			assert.NotNil(t, cli)

			assert.Len(t, cli.RecitalsSetSnapshot(), 0)
			assert.Len(t, cli.ChaptersSetSnapshot(), 0)
			assert.Len(t, cli.ArticlesSetSnapshot(), 0)
			assert.Len(t, cli.ArticleParagraphsSetSnapshot(), 0)
		})
	})

	t.Run("Given zero-value DataSettings (empty paths)", func(t *testing.T) {
		ds := &settings.DataSettings{}
		logger := zap.NewNop()

		t.Run("Should not error and produce empty datasets", func(t *testing.T) {
			cli, err := dal.NewGdprDataClient(ds, logger)

			assert.NoError(t, err)
			assert.NotNil(t, cli)
			assert.Empty(t, cli.RecitalsSetSnapshot())
			assert.Empty(t, cli.ChaptersSetSnapshot())
			assert.Empty(t, cli.ArticlesSetSnapshot())
			assert.Empty(t, cli.ArticleParagraphsSetSnapshot())
		})
	})

	t.Run("Given malformed recital JSON in data directory", func(t *testing.T) {
		ds := suite.badRecitalTempDataSettings(t)
		logger := zap.NewNop()

		t.Run("Should return an error and no client", func(t *testing.T) {
			cli, err := dal.NewGdprDataClient(ds, logger)

			assert.Error(t, err)
			assert.Nil(t, cli)
			assert.Contains(t, err.Error(), "recital missing ID")
		})
	})
}

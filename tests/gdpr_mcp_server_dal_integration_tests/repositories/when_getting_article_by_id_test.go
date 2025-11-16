package repositories_test

import (
	"testing"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal/repositories"
	"github.com/6022-labs/gdpr-mcp-server/tests/gdpr_mcp_server_dal_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenGettingArticleByIDTestingSuite struct {
	sut *repositories.ArticlesRepository

	gdprDataClientMock *gdpr_mcp_server_dal_mocks.MockGdprDataClientInterface
}

func WhenGettingArticleByIDBeforeEach(t *testing.T) *WhenGettingArticleByIDTestingSuite {
	mockController := gomock.NewController(t)

	gdprDataClientMock := gdpr_mcp_server_dal_mocks.NewMockGdprDataClientInterface(mockController)

	sut := repositories.NewArticlesRepository(gdprDataClientMock)

	return &WhenGettingArticleByIDTestingSuite{
		sut: sut,

		gdprDataClientMock: gdprDataClientMock,
	}
}

func TestWhenGettingArticleByID(t *testing.T) {
	t.Parallel()

	t.Run("Given a snapshot containing the article id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return the matching article and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleByIDBeforeEach(t)

			expected := &models.Article{ID: "art-1", Number: 1, Roman: "1", Title: "Subject-matter and objectives", NumberOfParagraphs: 2}
			snapshot := map[string]*models.Article{
				"art-1": expected,
			}
			suite.gdprDataClientMock.EXPECT().ArticlesSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetById("art-1")

			assert.NoError(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("Given a snapshot missing the article id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleByIDBeforeEach(t)

			snapshot := map[string]*models.Article{
				"art-2": {ID: "art-2", Number: 2, Roman: "2", Title: "Principles", NumberOfParagraphs: 6},
			}
			suite.gdprDataClientMock.EXPECT().ArticlesSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetById("art-1")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given a nil snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleByIDBeforeEach(t)
			suite.gdprDataClientMock.EXPECT().ArticlesSetSnapshot().Return(nil).Times(1)

			actual, err := suite.sut.GetById("art-1")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given an empty snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleByIDBeforeEach(t)
			suite.gdprDataClientMock.EXPECT().ArticlesSetSnapshot().Return(map[string]*models.Article{}).Times(1)

			actual, err := suite.sut.GetById("art-1")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given a non-empty snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil when id is empty", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleByIDBeforeEach(t)
			snapshot := map[string]*models.Article{
				"art-1": {ID: "art-1", Number: 1, Roman: "1", Title: "Subject-matter and objectives", NumberOfParagraphs: 2},
			}
			suite.gdprDataClientMock.EXPECT().ArticlesSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetById("")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})
}

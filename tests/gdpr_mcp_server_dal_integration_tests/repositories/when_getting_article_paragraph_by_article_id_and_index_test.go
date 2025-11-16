package repositories_test

import (
	"testing"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal/repositories"
	"github.com/6022-labs/gdpr-mcp-server/tests/gdpr_mcp_server_dal_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenGettingArticleParagraphByArticleIdAndIndexTestingSuite struct {
	sut *repositories.ArticleParagraphsRepository

	gdprDataClientMock *gdpr_mcp_server_dal_mocks.MockGdprDataClientInterface
}

func WhenGettingArticleParagraphByArticleIdAndIndexBeforeEach(t *testing.T) *WhenGettingArticleParagraphByArticleIdAndIndexTestingSuite {
	mockController := gomock.NewController(t)

	gdprDataClientMock := gdpr_mcp_server_dal_mocks.NewMockGdprDataClientInterface(mockController)

	sut := repositories.NewArticleParagraphsRepository(gdprDataClientMock)

	return &WhenGettingArticleParagraphByArticleIdAndIndexTestingSuite{
		sut: sut,

		gdprDataClientMock: gdprDataClientMock,
	}
}

func TestWhenGettingArticleParagraphByArticleIdAndIndex(t *testing.T) {
	t.Parallel()

	t.Run("Given a snapshot containing the article id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return the matching paragraph at index and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleParagraphByArticleIdAndIndexBeforeEach(t)

			p0 := &models.ArticleParagraph{Number: 1, ArticleId: "art-1", Texts: []string{"A"}}
			p1 := &models.ArticleParagraph{Number: 2, ArticleId: "art-1", Texts: []string{"B"}}
			snapshot := map[string][]*models.ArticleParagraph{
				"art-1": {p0, p1},
			}
			suite.gdprDataClientMock.EXPECT().ArticleParagraphsSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetByArticleIdAndIndex("art-1", 1)

			assert.NoError(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, p1, actual)
		})
	})

	t.Run("Given a snapshot missing the article id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleParagraphByArticleIdAndIndexBeforeEach(t)

			snapshot := map[string][]*models.ArticleParagraph{
				"art-2": {{Number: 1, ArticleId: "art-2", Texts: []string{"X"}}},
			}
			suite.gdprDataClientMock.EXPECT().ArticleParagraphsSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetByArticleIdAndIndex("art-1", 0)

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given a nil snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleParagraphByArticleIdAndIndexBeforeEach(t)
			suite.gdprDataClientMock.EXPECT().ArticleParagraphsSetSnapshot().Return(nil).Times(1)

			actual, err := suite.sut.GetByArticleIdAndIndex("art-1", 0)

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given an empty snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleParagraphByArticleIdAndIndexBeforeEach(t)
			suite.gdprDataClientMock.EXPECT().ArticleParagraphsSetSnapshot().Return(map[string][]*models.ArticleParagraph{}).Times(1)

			actual, err := suite.sut.GetByArticleIdAndIndex("art-1", 0)

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given a non-empty snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil when id is empty", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleParagraphByArticleIdAndIndexBeforeEach(t)
			snapshot := map[string][]*models.ArticleParagraph{
				"art-1": {{Number: 1, ArticleId: "art-1", Texts: []string{"A"}}},
			}
			suite.gdprDataClientMock.EXPECT().ArticleParagraphsSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetByArticleIdAndIndex("", 0)

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})

		t.Run("Should return error when index out of range", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleParagraphByArticleIdAndIndexBeforeEach(t)
			snapshot := map[string][]*models.ArticleParagraph{
				"art-1": {{Number: 1, ArticleId: "art-1", Texts: []string{"A"}}},
			}
			suite.gdprDataClientMock.EXPECT().ArticleParagraphsSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetByArticleIdAndIndex("art-1", 1)

			assert.Error(t, err)
			assert.EqualError(t, err, "index out of range")
			assert.Nil(t, actual)
		})

		t.Run("Should return error when slice empty", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingArticleParagraphByArticleIdAndIndexBeforeEach(t)
			snapshot := map[string][]*models.ArticleParagraph{
				"art-1": {},
			}
			suite.gdprDataClientMock.EXPECT().ArticleParagraphsSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetByArticleIdAndIndex("art-1", 0)

			assert.Error(t, err)
			assert.EqualError(t, err, "index out of range")
			assert.Nil(t, actual)
		})
	})
}

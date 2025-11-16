package repositories_test

import (
	"testing"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal/repositories"
	"github.com/6022-labs/gdpr-mcp-server/tests/gdpr_mcp_server_dal_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenGettingChapterByIDTestingSuite struct {
	sut *repositories.ChaptersRepository

	gdprDataClientMock *gdpr_mcp_server_dal_mocks.MockGdprDataClientInterface
}

func WhenGettingChapterByIDBeforeEach(t *testing.T) *WhenGettingChapterByIDTestingSuite {
	mockController := gomock.NewController(t)

	gdprDataClientMock := gdpr_mcp_server_dal_mocks.NewMockGdprDataClientInterface(mockController)

	sut := repositories.NewChaptersRepository(gdprDataClientMock)

	return &WhenGettingChapterByIDTestingSuite{
		sut: sut,

		gdprDataClientMock: gdprDataClientMock,
	}
}

func TestWhenGettingChapterByID(t *testing.T) {
	t.Parallel()

	t.Run("Given a snapshot containing the chapter id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return the matching chapter and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingChapterByIDBeforeEach(t)

			expected := &models.Chapter{ID: "ch-1", Roman: "I", Number: 1, Title: "General provisions", ArticlesIds: []string{"art-1", "art-2"}}
			snapshot := map[string]*models.Chapter{
				"ch-1": expected,
			}
			suite.gdprDataClientMock.EXPECT().ChaptersSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetById("ch-1")

			assert.NoError(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("Given a snapshot missing the chapter id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingChapterByIDBeforeEach(t)

			snapshot := map[string]*models.Chapter{
				"ch-2": {ID: "ch-2", Roman: "II", Number: 2, Title: "Principles", ArticlesIds: []string{"art-5"}},
			}
			suite.gdprDataClientMock.EXPECT().ChaptersSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetById("ch-1")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given a nil snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingChapterByIDBeforeEach(t)
			suite.gdprDataClientMock.EXPECT().ChaptersSetSnapshot().Return(nil).Times(1)

			actual, err := suite.sut.GetById("ch-1")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given an empty snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingChapterByIDBeforeEach(t)
			suite.gdprDataClientMock.EXPECT().ChaptersSetSnapshot().Return(map[string]*models.Chapter{}).Times(1)

			actual, err := suite.sut.GetById("ch-1")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given a non-empty snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil when id is empty", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingChapterByIDBeforeEach(t)
			snapshot := map[string]*models.Chapter{
				"ch-1": {ID: "ch-1", Roman: "I", Number: 1, Title: "General provisions"},
			}
			suite.gdprDataClientMock.EXPECT().ChaptersSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetById("")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})
}

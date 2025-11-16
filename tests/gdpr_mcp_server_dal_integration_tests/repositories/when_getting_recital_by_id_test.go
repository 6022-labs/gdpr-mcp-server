package repositories_test

import (
	"testing"

	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server/models"
	"github.com/6022-labs/gdpr-mcp-server/src/gdpr_mcp_server_dal/repositories"
	"github.com/6022-labs/gdpr-mcp-server/tests/gdpr_mcp_server_dal_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenGettingRecitalByIDTestingSuite struct {
	sut *repositories.RecitalsRepository

	gdprDataClientMock *gdpr_mcp_server_dal_mocks.MockGdprDataClientInterface
}

func WhenGettingRecitalByIDBeforeEach(t *testing.T) *WhenGettingRecitalByIDTestingSuite {
	mockController := gomock.NewController(t)

	gdprDataClientMock := gdpr_mcp_server_dal_mocks.NewMockGdprDataClientInterface(mockController)

	sut := repositories.NewRecitalsRepository(gdprDataClientMock)

	return &WhenGettingRecitalByIDTestingSuite{
		sut: sut,

		gdprDataClientMock: gdprDataClientMock,
	}
}

func TestWhenGettingRecitalByID(t *testing.T) {
	t.Parallel()

	t.Run("Given a snapshot containing the recital id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return the matching recital and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingRecitalByIDBeforeEach(t)

			expected := &models.Recital{ID: "rec-1", Number: 1, Texts: []string{"Lorem", "Ipsum"}}
			snapshot := map[string]*models.Recital{
				"rec-1": expected,
			}
			suite.gdprDataClientMock.EXPECT().RecitalsSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetById("rec-1")

			assert.NoError(t, err)
			assert.NotNil(t, actual)
			assert.Equal(t, expected, actual)
		})
	})

	t.Run("Given a snapshot missing the recital id", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingRecitalByIDBeforeEach(t)

			snapshot := map[string]*models.Recital{
				"rec-2": {ID: "rec-2", Number: 2, Texts: []string{"Other"}},
			}
			suite.gdprDataClientMock.EXPECT().RecitalsSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetById("rec-1")

			// Assert
			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given a nil snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingRecitalByIDBeforeEach(t)
			suite.gdprDataClientMock.EXPECT().RecitalsSetSnapshot().Return(nil).Times(1)

			actual, err := suite.sut.GetById("rec-1")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given an empty snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil result and no error", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingRecitalByIDBeforeEach(t)
			suite.gdprDataClientMock.EXPECT().RecitalsSetSnapshot().Return(map[string]*models.Recital{}).Times(1)

			actual, err := suite.sut.GetById("rec-1")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})

	t.Run("Given a non-empty snapshot", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return nil when id is empty", func(t *testing.T) {
			t.Parallel()

			suite := WhenGettingRecitalByIDBeforeEach(t)
			snapshot := map[string]*models.Recital{
				"rec-1": {ID: "rec-1", Number: 1, Texts: []string{"a"}},
			}
			suite.gdprDataClientMock.EXPECT().RecitalsSetSnapshot().Return(snapshot).Times(1)

			actual, err := suite.sut.GetById("")

			assert.NoError(t, err)
			assert.Nil(t, actual)
		})
	})
}

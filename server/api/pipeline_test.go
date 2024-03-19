package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/franela/goblin"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.woodpecker-ci.org/woodpecker/v2/server/model"
	"go.woodpecker-ci.org/woodpecker/v2/server/store/mocks"
)

func TestDeletePipelines(t *testing.T) {
	gin.SetMode(gin.TestMode)

	g := goblin.Goblin(t)
	g.Describe("Pipeline", func() {
		g.It("should parse pipeline filter", func() {
			pipelines := make([]*model.Pipeline, 1)

			mockStore := mocks.NewStore(t)
			mockStore.On("GetPipelineList", mock.Anything, mock.Anything, mock.Anything).Return(pipelines, nil)
			mockStore.On("DeletePipeline", mock.Anything).Return(nil)

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Set("store", mockStore)
			c.Request, _ = http.NewRequest("DELETE", "/?before=2023-01-16T15:00:00Z&after=2023-01-15T15:00:00Z", nil)

			DeletePipelines(c)

			assert.Equal(t, http.StatusNoContent, c.Writer.Status())
		})

		g.It("should parse pipeline filter with tz offset", func() {
			pipelines := make([]*model.Pipeline, 1)

			mockStore := mocks.NewStore(t)
			mockStore.On("GetPipelineList", mock.Anything, mock.Anything, mock.Anything).Return(pipelines, nil)
			mockStore.On("DeletePipeline", mock.Anything).Return(nil)

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Set("store", mockStore)
			c.Request, _ = http.NewRequest("DELETE", "/?before=2023-01-16T15:00:00%2B01:00&after=2023-01-15T15:00:00%2B01:00", nil)

			DeletePipelines(c)

			assert.Equal(t, http.StatusNoContent, c.Writer.Status())
		})

		g.It("should delete pipeline with filter", func() {
			pipelines := make([]*model.Pipeline, 1)

			mockStore := mocks.NewStore(t)
			mockStore.On("GetPipelineList", mock.Anything, mock.Anything, mock.Anything).Return(pipelines, nil)
			mockStore.On("DeletePipeline", mock.Anything).Return(nil)

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Set("store", mockStore)

			DeletePipelines(c)

			mockStore.AssertCalled(t, "GetPipelineList", mock.Anything, mock.Anything, mock.Anything)
			mockStore.AssertCalled(t, "DeletePipeline", mock.Anything)
			assert.Equal(t, http.StatusNoContent, c.Writer.Status())
		})

		g.It("should not delete pipeline", func() {
			pipelines := make([]*model.Pipeline, 0)

			mockStore := mocks.NewStore(t)
			mockStore.On("GetPipelineList", mock.Anything, mock.Anything, mock.Anything).Return(pipelines, nil)

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Set("store", mockStore)

			DeletePipelines(c)

			mockStore.AssertCalled(t, "GetPipelineList", mock.Anything, mock.Anything, mock.Anything)
			mockStore.AssertNotCalled(t, "DeletePipeline", mock.Anything)
			assert.Equal(t, http.StatusNoContent, c.Writer.Status())
		})

		g.It("should not parse pipeline filter", func() {
			mockStore := mocks.NewStore(t)

			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Set("store", mockStore)
			c.Request, _ = http.NewRequest("DELETE", "/?before=foo", nil)

			DeletePipelines(c)

			mockStore.AssertNotCalled(t, "GetPipelineList", mock.Anything, mock.Anything, mock.Anything)
			mockStore.AssertNotCalled(t, "DeletePipeline", mock.Anything)
			assert.Equal(t, http.StatusInternalServerError, c.Writer.Status())
		})
	})
}

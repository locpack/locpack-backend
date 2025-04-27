package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupControllerTest(t *testing.T, method, url string, body io.Reader) (*gin.Context, *httptest.ResponseRecorder) {
	t.Helper()
	gin.SetMode(gin.TestMode)

	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	request := httptest.NewRequest(method, url, body)

	if body != nil {
		bodyBytes, err := json.Marshal(body)
		assert.NoError(t, err)
		request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	ctx.Request = request

	return ctx, recorder
}

package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// Bad input is rejected before the database is touched, so a nil pool is fine here.
func TestInsertKeywordRejectsBadInput(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/keywords", InsertKeywordHandler(nil))

	for _, body := range []string{``, `{}`, `{"value": "   "}`, `not json`} {
		w := httptest.NewRecorder()
		r.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/keywords", strings.NewReader(body)))
		if w.Code != http.StatusBadRequest {
			t.Errorf("body %q: got %d, want 400", body, w.Code)
		}
	}
}

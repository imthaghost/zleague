package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"zleague/api/handlers"

	"github.com/matryer/is"
)

func TestHello(t *testing.T) {
	is := is.New(t)
	db := s.GetDB()
	m := s.GetManager()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := s.GetTestContext(req, rec)
	h := handlers.NewHandler(db, m)

	// make sure there is no error when we make a request
	is.NoErr(h.Hello(c))

	is.Equal(http.StatusOK, rec.Code)

	expect := `{"msg":"hello from zleague!"}`
	is.Equal(strings.TrimSpace(rec.Body.String()), expect)
}

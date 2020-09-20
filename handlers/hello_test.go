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
	c := s.GetContext(req, rec)
	h := handlers.New(db, m)

	is.NoErr(h.Hello(c))

	is.Equal(http.StatusOK, rec.Code)

	expect := `{"msg":"hello from zleague!"}`
	is.Equal(strings.TrimSpace(rec.Body.String()), expect)
}

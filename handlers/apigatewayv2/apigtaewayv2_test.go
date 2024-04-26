package apigatewayv2

import (
	"context"
	testengine "github.com/Drafteame/engine/test/engine"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewHandler(t *testing.T) {
	t.Run("should execute handler", func(t *testing.T) {
		s := http.NewServeMux()
		s.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		evt := HTTPRequest{
			RawPath: "/test",
			RequestContext: HTTPRequestContext{
				HTTP: HTTPRequestContextHTTPDescription{
					Method: "GET",
					Path:   "/test",
				},
			},
		}

		res, err := testengine.New(context.Background(), evt, NewHandler(s)).Run()

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}

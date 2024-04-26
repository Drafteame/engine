package apigatewayv1

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	testengine "github.com/Drafteame/engine/test/engine"
)

func TestNewHandler(t *testing.T) {
	t.Run("Should resolve handler", func(t *testing.T) {
		s := http.NewServeMux()
		s.HandleFunc("/test", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		evt := HTTPRequest{
			Path:       "/test",
			HTTPMethod: "GET",
		}

		res, err := testengine.New(context.TODO(), evt, NewHandler(s)).Run()

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}

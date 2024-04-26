package apigatewayv1

import (
	"context"
	testengine "github.com/Drafteame/engine/test/engine"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewHandler(t *testing.T) {
	t.Run("Should resolve handler", func(t *testing.T) {
		s := http.NewServeMux()
		s.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
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

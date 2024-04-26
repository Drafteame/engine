package apigatewayv2

import (
	"context"
	"github.com/Drafteame/engine"
	"github.com/Drafteame/engine/internal/request"
	"github.com/Drafteame/engine/internal/response"
	"net/http"
	"strings"
)

func NewHandler(handler http.Handler) engine.Handler[HTTPRequest, HTTPResponse] {
	return func(ctx context.Context, evt HTTPRequest) (HTTPResponse, error) {
		multiHeader := make(map[string][]string)
		for k, values := range evt.Headers {
			multiHeader[k] = strings.Split(values, ",")
		}

		req, err := request.New(ctx, request.Config{
			Path:        evt.RawPath,
			QueryString: evt.RawQueryString,
			Body:        evt.Body,
			IsBase64:    evt.IsBase64Encoded,
			Method:      evt.RequestContext.HTTP.Method,
			Context:     evt.RequestContext,
			SourceIP:    evt.RequestContext.HTTP.SourceIP,
			MultiHeader: multiHeader,
			Cookies:     evt.Cookies,
			RequestID:   evt.RequestContext.RequestID,
			Stage:       evt.RequestContext.Stage,
		})

		if err != nil {
			return HTTPResponse{}, err
		}

		res := response.New(new(HTTPResponse))

		handler.ServeHTTP(res, req)

		return *res.End(), nil
	}
}

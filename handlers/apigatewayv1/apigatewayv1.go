package apigatewayv1

import (
	"context"
	"errors"
	"github.com/Drafteame/engine"
	"github.com/Drafteame/engine/internal/request"
	"github.com/Drafteame/engine/internal/response"
	"net/http"
	"net/url"
)

var (
	ErrParsingPathFailed = errors.New("failed to parse path")
)

func NewHandler(handler http.Handler) engine.Handler[HTTPRequest, HTTPResponse] {
	return func(ctx context.Context, evt HTTPRequest) (HTTPResponse, error) {
		u, err := url.Parse(evt.Path)
		if err != nil {
			return HTTPResponse{}, errors.Join(err, ErrParsingPathFailed)
		}

		// querystring
		q := u.Query()
		for k, v := range evt.QueryStringParameters {
			q.Set(k, v)
		}

		for k, values := range evt.MultiValueQueryStringParameters {
			q[k] = values
		}

		req, err := request.New(ctx, request.Config{
			Path:        evt.Path,
			QueryString: q.Encode(),
			Body:        evt.Body,
			IsBase64:    evt.IsBase64Encoded,
			Method:      evt.HTTPMethod,
			Context:     evt.RequestContext,
			SourceIP:    evt.RequestContext.Identity.SourceIP,
			Headers:     evt.Headers,
			MultiHeader: evt.MultiValueHeaders,
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

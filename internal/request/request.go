package request

import (
	"context"
	"net/http"
)

type Config struct {
	Path        string
	QueryString string
	Body        string
	IsBase64    bool
	Method      string
	Context     any
	SourceIP    string
	Headers     map[string]string
	MultiHeader map[string][]string
	Cookies     []string
	RequestID   string
	Stage       string
}

func New(ctx context.Context, cfg Config) (*http.Request, error) {
	ri := requestInfo{
		path:        cfg.Path,
		queryString: cfg.QueryString,
		body:        cfg.Body,
		isBase64:    cfg.IsBase64,
		method:      cfg.Method,
		context:     cfg.Context,
		sourceIP:    cfg.SourceIP,
		headers:     cfg.Headers,
		multiHeader: cfg.MultiHeader,
		cookies:     cfg.Cookies,
		requestID:   cfg.RequestID,
		stage:       cfg.Stage,
	}

	return ri.toRequest(ctx)
}

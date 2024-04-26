package request

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type requestInfo struct {
	path        string
	queryString string
	body        string
	isBase64    bool
	method      string
	context     any
	sourceIP    string
	headers     map[string]string
	multiHeader map[string][]string
	cookies     []string
	requestID   string
	stage       string
}

func (ri requestInfo) toRequest(ctx context.Context) (*http.Request, error) {
	u, err := url.Parse(ri.path)
	if err != nil {
		return nil, errors.Join(err, ErrParsingPathFailed)
	}

	u.RawQuery = ri.queryString

	body, errBody := ri.decodeBody()
	if errBody != nil {
		return nil, errBody
	}

	req, errNew := http.NewRequest(ri.method, u.String(), strings.NewReader(body))
	if errNew != nil {
		return nil, errors.Join(errNew, ErrFailToCreateRequest)
	}

	// manually set RequestURI because NewRequest is for clients and req.RequestURI is for servers
	req.RequestURI = u.RequestURI()

	// remote addr
	req.RemoteAddr = ri.sourceIP

	// headers
	for k, v := range ri.headers {
		req.Header.Set(k, v)
	}

	// multi headers
	for k, values := range ri.multiHeader {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}

	// cookies
	for _, c := range ri.cookies {
		req.Header.Add("Cookie", c)
	}

	// content-length
	if req.Header.Get("Content-Length") == "" && body != "" {
		req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	}

	// custom fields
	req.Header.Set("X-Request-Id", ri.requestID)
	req.Header.Set("X-Stage", ri.stage)

	// custom context values
	req = req.WithContext(context.WithValue(ctx, "httpRequestContext", ri.context))

	// xray support
	if traceID := ctx.Value("x-amzn-trace-id"); traceID != nil {
		req.Header.Set("X-Amzn-Trace-Id", fmt.Sprintf("%v", traceID))
	}

	// host
	req.URL.Host = req.Header.Get("Host")
	req.Host = req.URL.Host

	return req, nil
}

func (ri requestInfo) decodeBody() (string, error) {
	if ri.isBase64 {
		b, errDecode := base64.StdEncoding.DecodeString(ri.body)
		if errDecode != nil {
			return "", errors.Join(errDecode, ErrDecodingBase64Body)
		}

		return string(b), nil
	}

	return ri.body, nil
}

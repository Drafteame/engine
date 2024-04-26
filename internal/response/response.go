package response

import (
	"bytes"
	"encoding/base64"
	"mime"
	"net/http"
	"strings"
)

type Out interface {
	SetStatusCode(int)
	SetHeaders(map[string]string)
	SetMultiValueHeaders(map[string][]string)
	SetBody(string)
	SetIsBase64Encoded(bool)
	SetCookies([]string)
}

// Writer implements the http.ResponseWriter interface
// in order to support the API Gateway Lambda HTTP "protocol".
type Writer[R Out] struct {
	out           R
	buf           bytes.Buffer
	header        http.Header
	wroteHeader   bool
	closeNotifyCh chan bool
}

// New returns a new response writer to capture http output.
func New[R Out](out R) *Writer[R] {
	return &Writer[R]{
		out:           out,
		closeNotifyCh: make(chan bool, 1),
	}
}

// Header implementation.
func (w *Writer[R]) Header() http.Header {
	if w.header == nil {
		w.header = make(http.Header)
	}

	return w.header
}

// Write implementation.
func (w *Writer[R]) Write(b []byte) (int, error) {
	if !w.wroteHeader {
		w.WriteHeader(http.StatusOK)
	}

	return w.buf.Write(b)
}

// WriteHeader implementation.
func (w *Writer[R]) WriteHeader(status int) {
	if w.wroteHeader {
		return
	}

	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "text/plain; charset=utf8")
	}

	w.out.SetStatusCode(status)

	h := make(map[string]string)
	mvh := make(map[string][]string)

	for k, v := range w.Header() {
		if len(v) == 1 {
			h[k] = v[0]
		} else if len(v) > 1 {
			mvh[k] = v
		}
	}

	w.out.SetHeaders(h)
	w.out.SetMultiValueHeaders(mvh)

	w.wroteHeader = true
}

// CloseNotify notify when the response is closed
func (w *Writer[R]) CloseNotify() <-chan bool {
	return w.closeNotifyCh
}

// End the request.
func (w *Writer[R]) End() R {
	isBin := isBinary(w.header)

	w.out.SetIsBase64Encoded(isBin)

	if isBin {
		w.out.SetBody(base64.StdEncoding.EncodeToString(w.buf.Bytes()))
	} else {
		w.out.SetBody(w.buf.String())
	}

	// see https://aws.amazon.com/blogs/compute/simply-serverless-using-aws-lambda-to-expose-custom-cookies-with-api-gateway/
	w.out.SetCookies(w.header["Set-Cookie"])
	w.header.Del("Set-Cookie")

	// notify end
	w.closeNotifyCh <- true

	return w.out
}

// isBinary returns true if the response represents binary.
func isBinary(h http.Header) bool {
	switch {
	case !isTextMime(h.Get("Content-Type")):
		return true
	case h.Get("Content-Encoding") == "gzip":
		return true
	default:
		return false
	}
}

// isTextMime returns true if the content type represents textual data.
func isTextMime(kind string) bool {
	mt, _, err := mime.ParseMediaType(kind)
	if err != nil {
		return false
	}

	if strings.HasPrefix(mt, "text/") {
		return true
	}

	switch mt {
	case "image/svg+xml", "application/json", "application/xml", "application/javascript":
		return true
	default:
		return false
	}
}

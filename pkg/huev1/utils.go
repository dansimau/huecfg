package huev1

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

// drainBody reads all of b to memory and then returns two equivalent
// ReadClosers yielding the same bytes.
// It returns an error if the initial slurp of all bytes fails. It does not attempt
// to make the returned ReadClosers have identical error-matching behavior.
// Copied from httputil
// https://golang.org/src/net/http/httputil/dump.go
func drainBody(b io.ReadCloser) (r1, r2 io.ReadCloser, err error) {
	if b == nil || b == http.NoBody {
		// No copying needed. Preserve the magic sentinel meaning of NoBody.
		return http.NoBody, http.NoBody, nil
	}
	var buf bytes.Buffer
	if _, err = buf.ReadFrom(b); err != nil {
		return nil, b, err
	}
	if err = b.Close(); err != nil {
		return nil, b, err
	}
	return ioutil.NopCloser(&buf), ioutil.NopCloser(bytes.NewReader(buf.Bytes())), nil
}

func readBody(resp *http.Response) ([]byte, error) {
	var err error
	var buf io.ReadCloser

	buf, resp.Body, err = drainBody(resp.Body)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(buf)
}

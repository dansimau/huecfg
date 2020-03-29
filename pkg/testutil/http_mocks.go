package testutil

import (
	"bufio"
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
)

// HTTPResponse is a struct representing the status code, headers and body of
// a HTTP response. This is constructed from a fixture file, which is a plain
// text file containing a HTTP response.
type HTTPResponse struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// URLToFixtureFile is a mapping of a URL path (to test) and the path to a file
// containing the raw HTTP response (headers and body).
type URLToFixtureFile struct {
	resp *HTTPResponse

	URLPath         string
	FixtureFilePath string
}

// HTTPResponse reads the fixture file and returns a HTTPResponse from the data
// in the file.
func (f *URLToFixtureFile) HTTPResponse() *HTTPResponse {
	if f.resp == nil {
		f.resp = httpResponseFromFile(f.FixtureFilePath)
	}
	return f.resp
}

// ServeMocksFromFile takes a list of URL-to-fixture file mappings and returns
// a listening HTTP test server that will respond on those paths with the
// contents of the file. Any unrecognised request paths will panic.
func ServeMocksFromFile(mocks ...*URLToFixtureFile) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		for _, mock := range mocks {
			if req.URL.EscapedPath() == mock.URLPath {
				resp := mock.HTTPResponse()

				for key, values := range resp.Headers {
					for _, val := range values {
						w.Header().Add(key, val)
					}
				}

				w.WriteHeader(resp.StatusCode)

				if _, err := w.Write(resp.Body); err != nil {
					panic(err)
				}
			}
		}
	}))
}

func httpResponseFromFile(path string) *HTTPResponse {
	r := &HTTPResponse{
		Headers: http.Header{},
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Read status code
	scanner.Scan()
	r.StatusCode = parseStatusCode(scanner.Bytes())

	// Read headers
	for scanner.Scan() {
		// empty newline means headers are finished
		if bytes.Equal(scanner.Bytes(), []byte{}) {
			break
		}

		key, values := parseHeader(scanner.Text())
		r.Headers[key] = append(r.Headers[key], values...)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Rest is the body
	buf := bytes.Buffer{}
	for scanner.Scan() {
		buf.Write(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	r.Body = buf.Bytes()

	return r
}

// parseStatusCode returns an int of the HTTP status code from the specified
// status line string. Panics if there is an error.
func parseStatusCode(statusLine []byte) (code int) {
	parts := bytes.SplitN(statusLine, []byte(" "), 3)

	code, err := strconv.Atoi(string(parts[1]))
	if err != nil {
		panic(err)
	}

	return code
}

// parseHeader returns a http.Header from the specified plain text line.
func parseHeader(text string) (key string, values []string) {
	parts := strings.SplitN(text, ":", 2)

	if len(parts) < 2 {
		panic("malformed input header")
	}

	key = parts[0]
	values = append(values, key)

	return key, values
}

package hue

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// Hue represents a Hue Bridge
type Hue struct {
	host     string
	username string

	Debug bool
	http  *http.Client

	Config *ConfigAPI
	Lights *LightsAPI
}

// NewConn creates a connection to a Hue Bridge.
func NewConn(host string, username string) *Hue {
	hue := &Hue{
		host:     host,
		username: username,

		http: &http.Client{},
	}

	hue.Config = &ConfigAPI{hue: hue}
	hue.Lights = &LightsAPI{hue: hue}

	return hue
}

func (h *Hue) debugReq(req *http.Request) {
	if h.Debug {
		println("\n--- DEBUG: HTTP request:")
		defer println("---")

		out, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "WARNING: Unable to dump request: %v\n", err)
			return
		}

		fmt.Fprint(os.Stderr, string(out))
	}
}

func (h *Hue) debugResp(resp *http.Response) {
	if h.Debug {
		println("\n--- DEBUG: HTTP response:")
		defer println("---")

		headers, err := httputil.DumpResponse(resp, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "WARNING: Unable to dump response header: %v\n", err)
			return
		} else {
			fmt.Fprint(os.Stderr, string(headers))
		}

		content, err := readBody(resp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "WARNING: Unable to read body: %v\n", err)
			return
		}

		if err := jsonutil.FprintBytes(os.Stderr, content); err != nil {
			fmt.Fprintf(os.Stderr, "WARNING: Unable to print JSON: %v\n", err)
			return
		}
	}
}

func (h *Hue) httpDo(req *http.Request) (*http.Response, error) {
	h.debugReq(req)
	resp, err := h.http.Do(req)
	h.debugResp(resp)
	return resp, err
}

func (h *Hue) httpGet(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return h.httpDo(req)
}

func (h *Hue) httpPost(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return h.httpDo(req)
}

func (h *Hue) requireUsername() (username string) {
	if h.username == "" {
		fmt.Fprintln(os.Stderr, "--username is required for this operation. See `huecfg api config create-user`.")
		os.Exit(1)
	}

	return h.username
}

// ResponseData holds the raw response from the Hue Bridge API response.
type ResponseData struct {
	bytes []byte
}

// Bytes returns the original (JSON) response from the Hue Bridge API.
func (r *ResponseData) Bytes() []byte {
	return r.bytes
}

package hue

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"

	"github.com/dansimau/huecfg/pkg/jsonutil"
)

// API represents the Hue Bridge API.
type API struct {
	Host     string
	Username string

	Debug bool
}

func (api *API) debugReq(req *http.Request) {
	if api.Debug {
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

func (api *API) debugResp(resp *http.Response) {
	if api.Debug {
		println("\n--- DEBUG: HTTP response:")
		defer println("---")

		headers, err := httputil.DumpResponse(resp, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "WARNING: Unable to dump response header: %v\n", err)
			return
		}

		fmt.Fprint(os.Stderr, string(headers))

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

func (api *API) httpDo(req *http.Request) (*http.Response, error) {
	api.debugReq(req)
	resp, err := http.DefaultClient.Do(req)
	api.debugResp(resp)
	return resp, err
}

func (api *API) httpGet(path string) (*http.Response, error) {
	url, err := api.url(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return api.httpDo(req)
}

func (api *API) httpPost(path string, body io.Reader) (*http.Response, error) {
	url, err := api.url(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return api.httpDo(req)
}

func (api *API) httpDelete(path string) (*http.Response, error) {
	url, err := api.url(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	return api.httpDo(req)
}

func (api *API) username() string {
	if api.Username == "" {
		return "nobody"
	}
	return api.Username
}

func (api *API) url(path string) (url string, err error) {
	urlParts := []string{}

	if !strings.HasPrefix(api.Host, "http://") && !strings.HasPrefix(api.Host, "https://") {
		urlParts = append(urlParts, "http://") // TODO: Use TLS by default
	}

	urlParts = append(urlParts, api.Host, path)

	return strings.Join(urlParts, ""), nil
}

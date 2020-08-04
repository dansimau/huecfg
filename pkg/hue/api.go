package hue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	Client *http.Client
	Debug  bool

	ctx context.Context
}

func (api *API) context() context.Context {
	if api.ctx == nil {
		return context.Background()
	}
	return api.ctx
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

func (api *API) get(path string) (response []byte, err error) {
	return api.httpReq("GET", path, nil)
}

func (api *API) delete(path string) (response []byte, err error) {
	return api.httpReq("DELETE", path, nil)
}

func (api *API) post(path string, data interface{}) (response []byte, err error) {
	return api.httpReq("POST", path, data)
}

func (api *API) put(path string, data interface{}) (response []byte, err error) {
	return api.httpReq("POST", path, data)
}

func (api *API) httpDo(req *http.Request) (*http.Response, error) {
	client := api.Client
	if client == nil {
		client = http.DefaultClient
	}

	api.debugReq(req)
	resp, err := client.Do(req)
	api.debugResp(resp)
	return resp, err
}

func (api *API) httpReq(method string, path string, data interface{}) (response []byte, err error) {
	url, err := api.url(path)
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(api.context(), method, url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := api.httpDo(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
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

func (api *API) WithContext(ctx context.Context) *API {
	newAPI := new(API)
	*newAPI = *api
	newAPI.ctx = ctx
	return newAPI
}

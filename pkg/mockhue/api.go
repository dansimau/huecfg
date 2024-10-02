package mockhue

import (
	"encoding/json"
	"net"
	"net/http"
	"sync"

	"github.com/dansimau/huecfg/pkg/huev1"
	"github.com/gorilla/mux"
	"github.com/mcuadros/go-lookup"
)

type Bridge struct {
	sync.RWMutex

	data   map[string]interface{}
	server *http.Server

	// We create and store the listener outside of the http server so that we
	// can call it directly to get the real listen address after a server is
	// started on an ephemeral port.
	socket net.Listener
}

func NewBridge(fixtureData []byte) (*Bridge, error) {
	b := &Bridge{}
	if err := b.SetData(fixtureData); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *Bridge) router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api", notImplementedHandler)
	r.HandleFunc("/api/{username}", notImplementedHandler)
	r.HandleFunc("/api/{username}/{entity}", b.entityGetHandler).Methods("GET")
	r.HandleFunc("/api/{username}/{entity}", b.entityDeleteHandler).Methods("DELETE")
	r.HandleFunc("/api/{username}/{entity}/{id}", b.entityGetHandler).Methods("GET")
	r.HandleFunc("/api/{username}/{entity}/{id}/{path:state}", b.entityGetHandler).Methods("GET")
	return r
}

// Close immediately closes connections and stops the listener.
func (b *Bridge) Close() {
	// Calling close on the http server will also close all listeners so
	// there's no need for us to do it explicitly.
	if b.server != nil {
		b.server.Close()
	}
}

// Start creates a server listening on an ephemeral port on localhost and
// returns the listening address.
func (b *Bridge) Start() (address string, err error) {
	if b.socket != nil {
		return b.socket.Addr().String(), nil
	}

	socket, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", err
	}

	b.socket = socket
	b.server = &http.Server{Handler: b.router()}

	// Serve always returns a non-nil error, e.g. it will eventually return
	// ErrServerClosed even when closed gracefully. There's not a great way of
	// trying to handle other errors here. But it's fine, because any fatal
	// error will surface on the client-side within a test.
	go b.server.Serve(b.socket)

	return b.socket.Addr().String(), nil
}

func notImplementedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (b *Bridge) entityDeleteHandler(w http.ResponseWriter, r *http.Request) {
}

func (b *Bridge) entityGetHandler(w http.ResponseWriter, r *http.Request) {
	jsonPath := []string{}
	vars := mux.Vars(r)

	jsonPath = append(jsonPath, vars["entity"])
	if id, ok := vars["id"]; ok {
		jsonPath = append(jsonPath, id)
	}
	if path, ok := vars["path"]; ok {
		jsonPath = append(jsonPath, path)
	}

	b.RLock()
	val, err := lookup.Lookup(b.data, jsonPath...)
	b.RUnlock()
	if err != nil {
		if err == lookup.ErrKeyNotFound {
			w.WriteHeader(404)

			res := huev1.StatusResponse{
				huev1.Status{
					Error: &huev1.Error{
						Address:     r.URL.Path,
						Description: "resource, " + r.URL.Path + ", not available",
						Type:        3,
					},
				},
			}

			resBytes, err := json.MarshalIndent(res, "", "  ")
			if err != nil {
				return
			}

			w.Write(resBytes)

		} else {
			w.WriteHeader(500)
		}
		w.Write([]byte(err.Error()))
		return
	}

	jsonBytes, err := json.Marshal(val.Interface())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(jsonBytes)
}

func (b *Bridge) SetData(jsonBytes []byte) error {
	var data map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return err
	}

	b.Lock()
	b.data = data
	b.Unlock()

	return nil
}

func (b *Bridge) SetLightData(jsonBytes []byte) error {
	var data interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return err
	}

	b.Lock()
	b.data["lights"] = data
	b.Unlock()

	return nil
}

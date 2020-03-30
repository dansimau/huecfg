package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// ConfigAPI is the API for config on the Hue Bridge.
type ConfigAPI struct {
	hue *Hue
}

// Config represents a config object as returned by the Hue Bridge API.
type Config struct {
	*responseData

	UTC        *AbsoluteTime
	APIVersion string
	Backup     struct {
		ErrorCode int
		Status    string
	}
	BridgeI          string
	DataStoreVersion string
	DHCP             bool
	FactoryNew       bool
	Gateway          string
	InternetServices struct {
		Internet     string
		RemoteAccess string
		SWUpdate     string
		Time         string
	}
	IPAddress        string
	LinkButton       bool
	MAC              string
	ModelID          string
	Name             string
	Netmask          string
	PortalConnection string
	PortalServices   bool
	PortalState      struct {
		Communication string
		Incoming      bool
		Outgoing      bool
		SignedOn      bool
	}
	ProxyAddress     string
	ProxyPort        int
	ReplacesBridgeID string
	StarterKitID     string
	SWUpdate         struct {
		CheckForUpdate bool
		DeviceTypes    struct {
			Bridge  bool
			Lights  []string
			Sensors []string
		}
		Notify      bool
		Text        string
		UpdateState int
		URL         string
	}
	SWUpdate2 struct {
		AutoInstall struct {
			On         bool
			UpdateTime string // TODO: Implement TimeInterval type
		}
		Bridge struct {
			LastInstall *AbsoluteTime
			State       string
		}
		CheckForUpdate bool
		LastChange     *AbsoluteTime
		State          string
	}
	SWVersion string
	Timezone  string
	Whitelist map[string]struct {
		CreateDate  *AbsoluteTime
		LastUseDate *AbsoluteTime
		Name        string
	}
	ZigbeeChannel int
}

// CreateUserResponse is the response from the user create API.
type CreateUserResponse struct {
	*responseData

	Success struct {
		Username string
	}
}

// CreateUser creates a new user. The link button on the bridge must be pressed and this command executed within 30
// seconds. Once a new user has been created, the user key is added to a 'whitelist', allowing access to API commands
// that require a whitelisted user. At present, all other API commands require a whitelisted user.
func (h *ConfigAPI) CreateUser(deviceType string, generateClientKey bool) (*CreateUserResponse, error) {
	params := struct {
		DeviceType        string `json:"devicetype"`
		GenerateClientKey *bool  `json:"generateclientkey,omitempty"`
	}{
		DeviceType: deviceType,
	}

	// An oddity of the hue bridge API: in testing, it accepted
	// generateclientkey: true but generateclientkey: false returned an error.
	// The field is marked as optional so it can be omitted.
	if generateClientKey {
		params.GenerateClientKey = &generateClientKey
	}

	postJSON, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	resp, err := h.hue.httpPost(fmt.Sprintf("%s/api", h.hue.host), bytes.NewBuffer(postJSON))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var obj *CreateUserResponse
	if err := json.Unmarshal(content, &obj); err != nil {
		return nil, err
	}
	obj.responseData = &responseData{content}

	return obj, nil
}

// Get returns list of all configuration elements in the bridge. Note all times are stored in UTC.
func (h *ConfigAPI) Get() (*Config, error) {
	resp, err := h.hue.httpGet(fmt.Sprintf("%s/api/%s/config", h.hue.host, h.hue.requireUsername()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var obj *Config
	if err := json.Unmarshal(content, &obj); err != nil {
		return nil, err
	}
	obj.responseData = &responseData{content}

	return obj, nil
}

package hue

import (
	"encoding/json"
)

// Config represents a config object as returned by the Hue Bridge API.
type Config struct {
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
		CreateDate  AbsoluteTime `json:"create date"`
		LastUseDate AbsoluteTime `json:"last use date"`
		Name        string
	}
	ZigbeeChannel int
}

// CreateUserResponse is the response from the user create API.
type CreateUserResponse struct {
	Success struct {
		Username string
	}
}

// CreateUser creates a new user. The link button on the bridge must be pressed and this command executed within 30
// seconds. Once a new user has been created, the user key is added to a 'whitelist', allowing access to API commands
// that require a whitelisted user. At present, all other API commands require a whitelisted user.
func (h *Hue) CreateUser(deviceType string, generateClientKey bool) (CreateUserResponse, error) {
	respBytes, err := h.API.CreateUser(deviceType, generateClientKey)
	if err != nil {
		return CreateUserResponse{}, err
	}

	var obj CreateUserResponse
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return CreateUserResponse{}, err
	}

	return obj, nil
}

// GetConfig returns list of all configuration elements in the bridge. Note all
// times are stored in UTC.
func (h *Hue) GetConfig() (Config, error) {
	respBytes, err := h.API.GetConfig()
	if err != nil {
		return Config{}, err
	}

	var obj Config
	if err := json.Unmarshal(respBytes, &obj); err != nil {
		return Config{}, err
	}

	return obj, nil
}

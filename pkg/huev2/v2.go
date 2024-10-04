package huev2

import "fmt"

func (api *API) GetResources() ([]byte, error) {
	return api.get("/clip/v2/resource")
}

// Lights
func (api *API) GetLights() ([]byte, error) {
	return api.get("/clip/v2/resource/light")
}

func (api *API) GetLight(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/light/%s", id))
}

func (api *API) PutLight(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/light/%s", id), data)
}

// Scene resource
func (api *API) GetScenes() ([]byte, error) {
	return api.get("/clip/v2/resource/scene")
}

func (api *API) GetScene(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/scene/%s", id))
}

func (api *API) PostScene(data interface{}) ([]byte, error) {
	return api.post("/clip/v2/resource/scene", data)
}

func (api *API) PutScene(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/scene/%s", id), data)
}

func (api *API) DeleteScene(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/scene/%s", id))
}

// Room resource
func (api *API) GetRooms() ([]byte, error) {
	return api.get("/clip/v2/resource/room")
}

func (api *API) GetRoom(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/room/%s", id))
}

func (api *API) PostRoom(data interface{}) ([]byte, error) {
	return api.post("/clip/v2/resource/room", data)
}

func (api *API) PutRoom(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/room/%s", id), data)
}

func (api *API) DeleteRoom(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/room/%s", id))
}

// Zone resource
func (api *API) GetZones() ([]byte, error) {
	return api.get("/clip/v2/resource/zone")
}

func (api *API) GetZone(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/zone/%s", id))
}

func (api *API) PostZone(data interface{}) ([]byte, error) {
	return api.post("/clip/v2/resource/zone", data)
}

func (api *API) PutZone(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/zone/%s", id), data)
}

func (api *API) DeleteZone(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/zone/%s", id))
}

// Bridge Home resource
func (api *API) GetBridgeHomes() ([]byte, error) {
	return api.get("/clip/v2/resource/bridge_home")
}

func (api *API) GetBridgeHome(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/bridge_home/%s", id))
}

// Grouped Light resource
func (api *API) GetGroupedLights() ([]byte, error) {
	return api.get("/clip/v2/resource/grouped_light")
}

func (api *API) GetGroupedLight(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/grouped_light/%s", id))
}

func (api *API) PutGroupedLight(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/grouped_light/%s", id), data)
}

// Device resource
func (api *API) GetDevices() ([]byte, error) {
	return api.get("/clip/v2/resource/device")
}

func (api *API) GetDevice(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/device/%s", id))
}

func (api *API) PutDevice(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/device/%s", id), data)
}

func (api *API) DeleteDevice(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/device/%s", id))
}

// Bridge resource
func (api *API) GetBridges() ([]byte, error) {
	return api.get("/clip/v2/resource/bridge")
}

func (api *API) GetBridge(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/bridge/%s", id))
}

func (api *API) PutBridge(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/bridge/%s", id), data)
}

// Device Software Update resource
func (api *API) GetDeviceSoftwareUpdates() ([]byte, error) {
	return api.get("/clip/v2/resource/device_software_update")
}

func (api *API) GetDeviceSoftwareUpdate(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/device_software_update/%s", id))
}

func (api *API) PutDeviceSoftwareUpdate(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/device_software_update/%s", id), data)
}

// Device Power resource
func (api *API) GetDevicePowers() ([]byte, error) {
	return api.get("/clip/v2/resource/device_power")
}

func (api *API) GetDevicePower(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/device_power/%s", id))
}

func (api *API) PutDevicePower(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/device_power/%s", id), data)
}

// Zigbee Connectivity resource
func (api *API) GetZigbeeConnectivities() ([]byte, error) {
	return api.get("/clip/v2/resource/zigbee_connectivity")
}

func (api *API) GetZigbeeConnectivity(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/zigbee_connectivity/%s", id))
}

func (api *API) PutZigbeeConnectivity(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/zigbee_connectivity/%s", id), data)
}

// ZGP Connectivity resource
func (api *API) GetZGPConnectivities() ([]byte, error) {
	return api.get("/clip/v2/resource/zgp_connectivity")
}

func (api *API) GetZGPConnectivity(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/zgp_connectivity/%s", id))
}

func (api *API) PutZGPConnectivity(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/zgp_connectivity/%s", id), data)
}

// Zigbee Device Discovery resource
func (api *API) GetZigbeeDeviceDiscoveries() ([]byte, error) {
	return api.get("/clip/v2/resource/zigbee_device_discovery")
}

func (api *API) GetZigbeeDeviceDiscovery(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/zigbee_device_discovery/%s", id))
}

func (api *API) PutZigbeeDeviceDiscovery(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/zigbee_device_discovery/%s", id), data)
}

// Motion resource
func (api *API) GetMotions() ([]byte, error) {
	return api.get("/clip/v2/resource/motion")
}

func (api *API) GetMotion(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/motion/%s", id))
}

func (api *API) PutMotion(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/motion/%s", id), data)
}

// Service Group resource
func (api *API) GetServiceGroups() ([]byte, error) {
	return api.get("/clip/v2/resource/service_group")
}

func (api *API) GetServiceGroup(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/service_group/%s", id))
}

func (api *API) PostServiceGroup(data interface{}) ([]byte, error) {
	return api.post("/clip/v2/resource/service_group", data)
}

func (api *API) PutServiceGroup(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/service_group/%s", id), data)
}

func (api *API) DeleteServiceGroup(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/service_group/%s", id))
}

// Grouped Motion resource
func (api *API) GetGroupedMotions() ([]byte, error) {
	return api.get("/clip/v2/resource/grouped_motion")
}

func (api *API) GetGroupedMotion(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/grouped_motion/%s", id))
}

func (api *API) PutGroupedMotion(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/grouped_motion/%s", id), data)
}

// Grouped Light Level resource
func (api *API) GetGroupedLightLevels() ([]byte, error) {
	return api.get("/clip/v2/resource/grouped_light_level")
}

func (api *API) GetGroupedLightLevel(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/grouped_light_level/%s", id))
}

func (api *API) PutGroupedLightLevel(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/grouped_light_level/%s", id), data)
}

// Camera Motion resource
func (api *API) GetCameraMotions() ([]byte, error) {
	return api.get("/clip/v2/resource/camera_motion")
}

func (api *API) GetCameraMotion(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/camera_motion/%s", id))
}

func (api *API) PutCameraMotion(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/camera_motion/%s", id), data)
}

// Temperature resource
func (api *API) GetTemperatures() ([]byte, error) {
	return api.get("/clip/v2/resource/temperature")
}

func (api *API) GetTemperature(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/temperature/%s", id))
}

func (api *API) PutTemperature(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/temperature/%s", id), data)
}

// Light Level resource
func (api *API) GetLightLevels() ([]byte, error) {
	return api.get("/clip/v2/resource/light_level")
}

func (api *API) GetLightLevel(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/light_level/%s", id))
}

func (api *API) PutLightLevel(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/light_level/%s", id), data)
}

// Button resource
func (api *API) GetButtons() ([]byte, error) {
	return api.get("/clip/v2/resource/button")
}

func (api *API) GetButton(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/button/%s", id))
}

func (api *API) PutButton(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/button/%s", id), data)
}

// Relative Rotary resource
func (api *API) GetRelativeRotaries() ([]byte, error) {
	return api.get("/clip/v2/resource/relative_rotary")
}

func (api *API) GetRelativeRotary(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/relative_rotary/%s", id))
}

func (api *API) PutRelativeRotary(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/relative_rotary/%s", id), data)
}

// Behavior Script resource
func (api *API) GetBehaviorScripts() ([]byte, error) {
	return api.get("/clip/v2/resource/behavior_script")
}

func (api *API) GetBehaviorScript(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/behavior_script/%s", id))
}

// Behavior Instance resource
func (api *API) GetBehaviorInstances() ([]byte, error) {
	return api.get("/clip/v2/resource/behavior_instance")
}

func (api *API) GetBehaviorInstance(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/behavior_instance/%s", id))
}

func (api *API) PostBehaviorInstance(data interface{}) ([]byte, error) {
	return api.post("/clip/v2/resource/behavior_instance", data)
}

func (api *API) PutBehaviorInstance(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/behavior_instance/%s", id), data)
}

func (api *API) DeleteBehaviorInstance(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/behavior_instance/%s", id))
}

// Geofence Client resource
func (api *API) GetGeofenceClients() ([]byte, error) {
	return api.get("/clip/v2/resource/geofence_client")
}

func (api *API) GetGeofenceClient(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/geofence_client/%s", id))
}

func (api *API) PostGeofenceClient(data interface{}) ([]byte, error) {
	return api.post("/clip/v2/resource/geofence_client", data)
}

func (api *API) PutGeofenceClient(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/geofence_client/%s", id), data)
}

func (api *API) DeleteGeofenceClient(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/geofence_client/%s", id))
}

// Geolocation resource
func (api *API) GetGeolocations() ([]byte, error) {
	return api.get("/clip/v2/resource/geolocation")
}

func (api *API) GetGeolocation(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/geolocation/%s", id))
}

func (api *API) PutGeolocation(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/geolocation/%s", id), data)
}

// Entertainment Configuration resource
func (api *API) GetEntertainmentConfigurations() ([]byte, error) {
	return api.get("/clip/v2/resource/entertainment_configuration")
}

func (api *API) GetEntertainmentConfiguration(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/entertainment_configuration/%s", id))
}

func (api *API) PostEntertainmentConfiguration(data interface{}) ([]byte, error) {
	return api.post("/clip/v2/resource/entertainment_configuration", data)
}

func (api *API) PutEntertainmentConfiguration(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/entertainment_configuration/%s", id), data)
}

func (api *API) DeleteEntertainmentConfiguration(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/entertainment_configuration/%s", id))
}

// Entertainment resource
func (api *API) GetEntertainments() ([]byte, error) {
	return api.get("/clip/v2/resource/entertainment")
}

func (api *API) GetEntertainment(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/entertainment/%s", id))
}

func (api *API) PutEntertainment(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/entertainment/%s", id), data)
}

// Homekit resource
func (api *API) GetHomekits() ([]byte, error) {
	return api.get("/clip/v2/resource/homekit")
}

func (api *API) GetHomekit(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/homekit/%s", id))
}

func (api *API) PutHomekit(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/homekit/%s", id), data)
}

// Matter resource
func (api *API) GetMatters() ([]byte, error) {
	return api.get("/clip/v2/resource/matter")
}

func (api *API) GetMatter(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/matter/%s", id))
}

func (api *API) PutMatter(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/matter/%s", id), data)
}

// Matter Fabric resource
func (api *API) GetMatterFabrics() ([]byte, error) {
	return api.get("/clip/v2/resource/matter_fabric")
}

func (api *API) GetMatterFabric(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/matter_fabric/%s", id))
}

func (api *API) DeleteMatterFabric(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/matter_fabric/%s", id))
}

// Smart Scene resource
func (api *API) GetSmartScenes() ([]byte, error) {
	return api.get("/clip/v2/resource/smart_scene")
}

func (api *API) GetSmartScene(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/smart_scene/%s", id))
}

func (api *API) PostSmartScene(data interface{}) ([]byte, error) {
	return api.post("/clip/v2/resource/smart_scene", data)
}

func (api *API) PutSmartScene(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/smart_scene/%s", id), data)
}

func (api *API) DeleteSmartScene(id string) ([]byte, error) {
	return api.delete(fmt.Sprintf("/clip/v2/resource/smart_scene/%s", id))
}

// Contact resource
func (api *API) GetContacts() ([]byte, error) {
	return api.get("/clip/v2/resource/contact")
}

func (api *API) GetContact(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/contact/%s", id))
}

func (api *API) PutContact(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/contact/%s", id), data)
}

// Tamper resource
func (api *API) GetTampers() ([]byte, error) {
	return api.get("/clip/v2/resource/tamper")
}

func (api *API) GetTamper(id string) ([]byte, error) {
	return api.get(fmt.Sprintf("/clip/v2/resource/tamper/%s", id))
}

func (api *API) PutTamper(id string, data interface{}) ([]byte, error) {
	return api.put(fmt.Sprintf("/clip/v2/resource/tamper/%s", id), data)
}

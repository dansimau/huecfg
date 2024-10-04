package cmd

type apiV2Cmd struct {
	BehaviorInstance           *apiV2BehaviorInstanceCmd           `command:"behavior-instance" description:"API to manage instances of script"`
	BehaviorScript             *apiV2BehaviorScriptCmd             `command:"behavior-script" description:"API to discover available scripts that can be instantiated"`
	Bridge                     *apiV2BridgeCmd                     `command:"bridge" description:"API to manage the bridge"`
	BridgeHome                 *apiV2BridgeHomeCmd                 `command:"bridge-home" description:"API to manage bridge homes"`
	Button                     *apiV2ButtonCmd                     `command:"button" description:"API to manage button services"`
	CameraMotion               *apiV2CameraMotionCmd               `command:"camera-motion" description:"API to manage camera_motion services"`
	Contact                    *apiV2ContactCmd                    `command:"contact" description:"API to manage contact sensor state"`
	Device                     *apiV2DeviceCmd                     `command:"device" description:"API to manage devices"`
	DevicePower                *apiV2DevicePowerCmd                `command:"device-power" description:"API to manage device power services"`
	DeviceSoftwareUpdate       *apiV2DeviceSoftwareUpdateCmd       `command:"device-software-update" description:"API to manage device update services"`
	Entertainment              *apiV2EntertainmentCmd              `command:"entertainment" description:"API to manage entertainment services"`
	EntertainmentConfiguration *apiV2EntertainmentConfigurationCmd `command:"entertainment-configuration" description:"API to manage entertainment configurations"`
	GeofenceClient             *apiV2GeofenceClientCmd             `command:"geofence-client" description:"API for geofencing functionality"`
	Geolocation                *apiV2GeolocationCmd                `command:"geolocation" description:"API for setting the geolocation"`
	GroupedLight               *apiV2GroupedLightCmd               `command:"grouped-light" description:"API to manage grouped light services"`
	GroupedLightLevel          *apiV2GroupedLightLevelCmd          `command:"grouped-light-level" description:"API to manage grouped light-level services"`
	GroupedMotion              *apiV2GroupedMotionCmd              `command:"grouped-motion" description:"API to manage grouped motion services"`
	Homekit                    *apiV2HomekitCmd                    `command:"homekit" description:"API to manage homekit service"`
	Light                      *apiV2LightCmd                      `command:"light" description:"API to manage light services"`
	LightLevel                 *apiV2LightLevelCmd                 `command:"light-level" description:"API to manage light level services"`
	Matter                     *apiV2MatterCmd                     `command:"matter" description:"API to manage matter service"`
	MatterFabric               *apiV2MatterFabricCmd               `command:"matter-fabric" description:"API to manage matter fabrics"`
	Motion                     *apiV2MotionCmd                     `command:"motion" description:"API to manage motion services"`
	RelativeRotary             *apiV2RelativeRotaryCmd             `command:"relative-rotary" description:"API to manage relative rotary services"`
	Resource                   *apiV2ResourceCmd                   `command:"resource" description:"API to retrieve all resources"`
	Room                       *apiV2RoomCmd                       `command:"room" description:"API to manage rooms"`
	Scene                      *apiV2SceneCmd                      `command:"scene" description:"API to manage scenes"`
	ServiceGroup               *apiV2ServiceGroupCmd               `command:"service-group" description:"API to manage service group services"`
	SmartScene                 *apiV2SmartSceneCmd                 `command:"smart-scene" description:"API to manage smart scenes"`
	Tamper                     *apiV2TamperCmd                     `command:"tamper" description:"API to manage device tamper state"`
	Temperature                *apiV2TemperatureCmd                `command:"temperature" description:"API to manage temperature services"`
	ZgpConnectivity            *apiV2ZGPConnectivityCmd            `command:"zgp-connectivity" description:"API to manage zgp connectivity services"`
	ZigbeeConnectivity         *apiV2ZigbeeConnectivityCmd         `command:"zigbee-connectivity" description:"API to manage zigbee connectivity services"`
	ZigbeeDeviceDiscovery      *apiV2ZigbeeDeviceDiscoveryCmd      `command:"zigbee-device-discovery" description:"API to manage zigbee device discovery service"`
	Zone                       *apiV2ZoneCmd                       `command:"zone" description:"API to manage zones"`
}

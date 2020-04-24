huecfg
======

hucfg is a command-line interface and configuration tool for managing lights
on a Philips Hue Bridge.

**Key features:**

* Configure lights, switches and other devices from a YAML file, allowing you
  to reset your bridge and rebuild from config. Also unlocks additional
  functionality of the bridge not available through the Hue app.
* CLI to read and set light state, color, brightness, etc.
* CLI to interact directly with the Hue Bridge's HTTP/JSON API.

Getting started
---------------

TODO

`pkg/hue` Supported APIs
------------------------

| API               | Read | Search | Create | Set |
|-------------------|------|--------|--------|-----|
| Lights            | [x]  |  [ ]   |  N/A   | [ ] |
| Groups            | [x]  |  N/A   |  [ ]   | [ ] |
| Schedules         | [x]  |  N/A   |  [ ]   | [ ] |
| Scenes            | [x]  |  N/A   |  [ ]   | [ ] |
| Sensors           | [x]  |  [ ]   |  N/A   | [ ] |
| Rules             | [x]  |  N/A   |  [ ]   | [ ] |
| Configuration     | [x]¹ |  N/A   |  N/A   | [ ] |
| Info (deprecated) |  -   |   -    |   -    |  -  |
| Resource links    | [x]  |  N/A   |  [ ]   | [ ] |
| Capabilities      | [ ]  |  N/A   |  N/A   | N/A |

¹ "Get full state" not implemented

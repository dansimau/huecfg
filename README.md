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

**Key:**

* ✅ = supported
* 🚧 = under construction / planned
* ❌ = no plans to support
* - = not applicable

| API               | List | Show | Search | Create | Set |
|-------------------|------|------|--------|--------|-----|
| Lights            |  ✅  |  ✅  |   🚧   |   -    |  🚧 |
| Groups            |  ✅  |  ✅  |   -    |   🚧   |  🚧 |
| Schedules         |  ✅  |  ✅  |   -    |   🚧   |  🚧 |
| Scenes            |  ✅  |  ✅  |   -    |   🚧   |  🚧 |
| Sensors           |  ✅  |  ✅  |   🚧   |   -    |  🚧 |
| Rules             |  ✅  |  ✅  |   -    |   🚧   |  🚧 |
| Configuration     |  ✅¹ |  -   |   -    |   -    |  🚧 |
| Info (deprecated) |  ❌  |  ❌  |   ❌   |   ❌   |  ❌ |
| Resource links    |  ✅  |  ✅  |   -    |   🚧   |  🚧 |
| Capabilities      |  ✅  |  -   |   -    |   -    |  -  |

¹ "Get full state" not implemented

module github.com/dansimau/huecfg

go 1.14

require (
	github.com/alecthomas/chroma v0.7.2
	github.com/davecgh/go-spew v1.1.1
	github.com/iancoleman/strcase v0.0.0-20191112232945-16388991a334
	github.com/jessevdk/go-flags v1.4.0
	github.com/mcuadros/go-lookup v0.0.0-20200330054200-b4062b0c4c85
	github.com/mikefarah/yq/v3 v3.0.0-20200501003153-6fc3566acd3a
	github.com/mitchellh/go-homedir v1.1.0
	github.com/olekukonko/tablewriter v0.0.4
	github.com/stretchr/testify v1.5.1
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/mcuadros/go-lookup => ../go-lookup

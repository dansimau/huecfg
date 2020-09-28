package hue_test

import (
	"testing"

	"github.com/dansimau/huecfg/pkg/fixtures"
	"github.com/dansimau/huecfg/pkg/hue"
	"github.com/dansimau/huecfg/pkg/mockhue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetLights(t *testing.T) {
	test(t, func(client *hue.Hue, mockBridge *mockhue.Bridge) {
		lights, err := client.GetLights()
		require.NoError(t, err)

		assert.Equal(t, []hue.Light{
			fixtures.DefaultHueColorLamp7,
		}, lights)
	})
}

func TestGetLight(t *testing.T) {
	test(t, func(client *hue.Hue, mockBridge *mockhue.Bridge) {
		light, err := client.GetLight("1")
		require.NoError(t, err)

		assert.Equal(t, fixtures.DefaultHueColorLamp7, light)
	})
}

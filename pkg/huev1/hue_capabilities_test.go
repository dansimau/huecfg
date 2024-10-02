package huev1_test

import (
	"testing"

	"github.com/dansimau/huecfg/pkg/fixtures"
	"github.com/dansimau/huecfg/pkg/huev1"
	"github.com/dansimau/huecfg/pkg/mockhue"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCapabilities(t *testing.T) {
	test(t, func(client *huev1.Hue, mockBridge *mockhue.Bridge) {
		capabilities, err := client.GetCapabilities()
		require.NoError(t, err)

		assert.Equal(t, fixtures.Capabilities, capabilities)
	})
}

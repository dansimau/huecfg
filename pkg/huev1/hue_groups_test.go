package huev1_test

import (
	"testing"

	"github.com/dansimau/huecfg/pkg/huev1"
	"github.com/dansimau/huecfg/pkg/mockhue"
)

func TestGetGroupEmptyID(t *testing.T) {
	test(t, func(client *huev1.Hue, mockBridge *mockhue.Bridge) {
		testErrorEmptyID(t, client.GetGroup)
	})
}

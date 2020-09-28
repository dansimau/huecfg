package hue_test

import (
	"reflect"
	"testing"

	"github.com/dansimau/huecfg/pkg/hue"
	"github.com/dansimau/huecfg/pkg/mockhue"
	"github.com/stretchr/testify/assert"
)

func TestEmptyIDs(t *testing.T) {
	test(t, func(client *hue.Hue, mockBridge *mockhue.Bridge) {
		// List of API methods to test
		fnsToTest := []interface{}{
			client.GetGroup,
			client.GetLight,
			client.GetResourceLink,
			client.GetRule,
			client.GetScene,
			client.GetSchedule,
			client.GetSensor,
			client.DeleteLight,
		}

		for _, fn := range fnsToTest {
			t.Run(funcName(fn), func(t *testing.T) {
				v := reflect.ValueOf(fn)
				// call each method above, with an empty string as the ID
				results := v.Call([]reflect.Value{
					reflect.ValueOf(""),
				})

				err := results[1].Interface().(error)

				assert.EqualError(t, hue.ErrEmptyID, err.Error())
			})
		}
	})
}

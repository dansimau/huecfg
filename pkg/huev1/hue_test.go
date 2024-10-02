package huev1_test

import (
	"reflect"
	"testing"

	"github.com/dansimau/huecfg/pkg/huev1"
	"github.com/dansimau/huecfg/pkg/mockhue"
	"github.com/stretchr/testify/assert"
)

func testErrorEmptyID(t *testing.T, fn interface{}) {
	// call fn with empty string as the argument
	results := reflect.ValueOf(fn).Call([]reflect.Value{
		reflect.ValueOf(""),
	})

	err := results[1].Interface().(error)

	assert.EqualError(t, huev1.ErrEmptyID, err.Error())
}

func TestErrorEmptyIDs(t *testing.T) {
	test(t, func(client *huev1.Hue, mockBridge *mockhue.Bridge) {
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
				testErrorEmptyID(t, fn)
			})
		}
	})
}

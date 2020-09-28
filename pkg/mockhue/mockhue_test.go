package mockhue_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/dansimau/huecfg/pkg/mockhue"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMockHue(t *testing.T) {
	mockBridge, err := mockhue.NewBridge([]byte(`
	{
		"lights": {
			"foo": {
				"state": {
					"foo": "bar"
				}
			}
		}
	}
	`))
	require.NoError(t, err)
	defer mockBridge.Close()

	address, err := mockBridge.Start()
	require.NoError(t, err)

	resp, err := http.Get("http://" + address + "/api/foouser/lights/foo/state")
	require.NoError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	assert.Equal(t, `{"foo":"bar"}`, string(body))
}

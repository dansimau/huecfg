package huev1_test

import (
	"testing"

	"github.com/dansimau/huecfg/pkg/huev1"
	"github.com/dansimau/huecfg/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIGetLights(t *testing.T) {
	t.Parallel()

	mock := &testutil.URLToFixtureFile{
		URLPath:         "/api/test/lights",
		FixtureFilePath: "../../test/fixtures/generic/lights-getall.txt",
	}
	server := testutil.ServeMocksFromFile(mock)
	defer server.Close()

	api := huev1.API{
		Host:     server.URL,
		Username: "test",
	}
	resp, err := api.GetLights()
	require.NoError(t, err)

	assert.Equal(t, mock.HTTPResponse().Body, resp)
}

func TestAPIGetLight(t *testing.T) {
	t.Parallel()

	mock := &testutil.URLToFixtureFile{
		URLPath:         "/api/test/lights/1",
		FixtureFilePath: "../../test/fixtures/generic/lights-getlight.txt",
	}
	server := testutil.ServeMocksFromFile(mock)
	defer server.Close()

	api := huev1.API{
		Host:     server.URL,
		Username: "test",
	}
	resp, err := api.GetLight("1")
	require.NoError(t, err)

	assert.Equal(t, mock.HTTPResponse().Body, resp)
}

func TestAPIGetLightErr(t *testing.T) {
	t.Parallel()

	mock := &testutil.URLToFixtureFile{
		URLPath:         "/api/test/lights/2",
		FixtureFilePath: "../../test/fixtures/generic/lights-getlight-err.txt",
	}
	server := testutil.ServeMocksFromFile(mock)
	defer server.Close()

	api := huev1.API{
		Host:     server.URL,
		Username: "test",
	}
	resp, err := api.GetLight("2")
	require.NoError(t, err)

	assert.Equal(t, mock.HTTPResponse().Body, resp)
}

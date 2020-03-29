package cmd_test

import (
	"testing"

	"github.com/dansimau/huecfg/pkg/hue"
	"github.com/dansimau/huecfg/pkg/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLights(t *testing.T) {
	mock := &testutil.URLToFixtureFile{
		URLPath:         "/api/test/lights",
		FixtureFilePath: "fixtures/generic/lights.txt",
	}
	server := testutil.ServeMocksFromFile(mock)
	defer server.Close()

	hue := hue.NewConn(server.URL, "test")
	resp, err := hue.Lights.GetAll()
	require.NoError(t, err)

	assert.Equal(t, len(mock.HTTPResponse().Body), len(resp.ResponseData.Bytes()))
}

package hue_test

import (
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/dansimau/huecfg/pkg/fixtures"
	"github.com/dansimau/huecfg/pkg/hue"
	"github.com/dansimau/huecfg/pkg/mockhue"
	"github.com/stretchr/testify/require"
)

type testFunc func(client *hue.Hue, mockBridge *mockhue.Bridge)

func createDefaultClientServer(t *testing.T) (*hue.Hue, *mockhue.Bridge) {
	mockBridge, err := mockhue.NewBridge(fixtures.Default)
	require.NoError(t, err)

	address, err := mockBridge.Start()
	require.NoError(t, err)

	return hue.NewConn(address, "test"), mockBridge
}

func funcName(i interface{}) string {
	fullFuncName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(fullFuncName, ".")
	funcName := parts[len(parts)-1]
	return funcName
}

func test(t *testing.T, f testFunc) {
	t.Parallel()

	client, mockBridge := createDefaultClientServer(t)
	defer mockBridge.Close()

	f(client, mockBridge)
}

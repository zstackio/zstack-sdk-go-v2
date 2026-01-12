// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2Gateway(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2Gateway(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2Gateway error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2Gateway result count: %d", len(result))
}

func TestUpdateBareMetal2Gateway(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2Gateway(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2Gateway Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2Gateway found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBareMetal2GatewayParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBareMetal2GatewayParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBareMetal2Gateway(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2Gateway error: %v", err)
		return
	}
	golog.Infof("UpdateBareMetal2Gateway result: %s", result.UUID)
}

func TestDeleteBareMetal2Gateway(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteBareMetal2Gateway is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2Gateway(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteBareMetal2Gateway Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2Gateway found to test Delete")
		return
	}

	err = accountLoginCli.DeleteBareMetal2Gateway(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteBareMetal2Gateway error: %v", err)
		return
	}
	golog.Infof("DeleteBareMetal2Gateway succeeded for UUID: %s", list[0].UUID)
}

func TestReconnectBareMetal2Gateway(t *testing.T) {
	// ReconnectBareMetal2Gateway operation
	t.Skip("TestReconnectBareMetal2Gateway requires manual implementation")

}

func TestAddBareMetal2Gateway(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddBareMetal2Gateway requires valid creation parameters")

}

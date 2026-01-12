// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPciDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPciDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPciDevice error: %v", err)
		return
	}
	golog.Infof("QueryPciDevice result count: %d", len(result))
}

func TestUpdatePciDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPciDevice(&queryParam)
	if err != nil {
		t.Errorf("TestUpdatePciDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PciDevice found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdatePciDeviceParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdatePciDeviceParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdatePciDevice(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdatePciDevice error: %v", err)
		return
	}
	golog.Infof("UpdatePciDevice result: %s", result.UUID)
}

func TestDeletePciDevice(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePciDevice is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPciDevice(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePciDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PciDevice found to test Delete")
		return
	}

	err = accountLoginCli.DeletePciDevice(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePciDevice error: %v", err)
		return
	}
	golog.Infof("DeletePciDevice succeeded for UUID: %s", list[0].UUID)
}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryScsiLun(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryScsiLun(&queryParam)
	if err != nil {
		t.Errorf("TestQueryScsiLun error: %v", err)
		return
	}
	golog.Infof("QueryScsiLun result count: %d", len(result))
}
func TestGetScsiLun(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryScsiLun(&queryParam)
	if err != nil {
		t.Errorf("TestGetScsiLun Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ScsiLun found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetScsiLun(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetScsiLun error: %v", err)
		return
	}
	golog.Infof("GetScsiLun result: %s", result.UUID)
}

func TestUpdateScsiLun(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryScsiLun(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateScsiLun Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ScsiLun found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateScsiLunParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateScsiLunParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateScsiLun(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateScsiLun error: %v", err)
		return
	}
	golog.Infof("UpdateScsiLun result: %s", result.UUID)
}

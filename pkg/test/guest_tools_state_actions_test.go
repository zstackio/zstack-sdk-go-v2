// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGuestToolsState(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGuestToolsState(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGuestToolsState error: %v", err)
		return
	}
	golog.Infof("QueryGuestToolsState result count: %d", len(result))
}

func TestUpdateGuestToolsState(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryGuestToolsState(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateGuestToolsState Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No GuestToolsState found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateGuestToolsStateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateGuestToolsStateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateGuestToolsState(*list[0].VmInstanceUuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateGuestToolsState error: %v", err)
		return
	}
	golog.Infof("UpdateGuestToolsState result: %s", result.VmInstanceUuid)
}

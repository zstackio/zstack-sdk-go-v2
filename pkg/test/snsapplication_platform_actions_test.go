// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSApplicationPlatform(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSApplicationPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSApplicationPlatform error: %v", err)
		return
	}
	golog.Infof("QuerySNSApplicationPlatform result count: %d", len(result))
}
func TestGetSNSApplicationPlatform(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSApplicationPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSApplicationPlatform Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSApplicationPlatform found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSApplicationPlatform(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSApplicationPlatform error: %v", err)
		return
	}
	golog.Infof("GetSNSApplicationPlatform result: %s", result.UUID)
}

func TestUpdateSNSApplicationPlatform(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSApplicationPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSApplicationPlatform Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSApplicationPlatform found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSApplicationPlatformParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSApplicationPlatformParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSApplicationPlatform(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSApplicationPlatform error: %v", err)
		return
	}
	golog.Infof("UpdateSNSApplicationPlatform result: %s", result.UUID)
}

func TestDeleteSNSApplicationPlatform(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSNSApplicationPlatform is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSApplicationPlatform(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSNSApplicationPlatform Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSApplicationPlatform found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSNSApplicationPlatform(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSNSApplicationPlatform error: %v", err)
		return
	}
	golog.Infof("DeleteSNSApplicationPlatform succeeded for UUID: %s", list[0].UUID)
}

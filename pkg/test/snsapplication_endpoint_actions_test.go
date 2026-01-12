// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSApplicationEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSApplicationEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSApplicationEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSApplicationEndpoint result count: %d", len(result))
}
func TestGetSNSApplicationEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSApplicationEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSApplicationEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSApplicationEndpoint found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSApplicationEndpoint(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSApplicationEndpoint error: %v", err)
		return
	}
	golog.Infof("GetSNSApplicationEndpoint result: %s", result.UUID)
}

func TestUpdateSNSApplicationEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSApplicationEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSApplicationEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSApplicationEndpoint found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSApplicationEndpointParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSApplicationEndpointParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSApplicationEndpoint(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSApplicationEndpoint error: %v", err)
		return
	}
	golog.Infof("UpdateSNSApplicationEndpoint result: %s", result.UUID)
}

func TestDeleteSNSApplicationEndpoint(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSNSApplicationEndpoint is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSApplicationEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSNSApplicationEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSApplicationEndpoint found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSNSApplicationEndpoint(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSNSApplicationEndpoint error: %v", err)
		return
	}
	golog.Infof("DeleteSNSApplicationEndpoint succeeded for UUID: %s", list[0].UUID)
}

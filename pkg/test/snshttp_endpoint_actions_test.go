// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSHttpEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSHttpEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSHttpEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSHttpEndpoint result count: %d", len(result))
}
func TestGetSNSHttpEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSHttpEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSHttpEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSHttpEndpoint found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSHttpEndpoint(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSHttpEndpoint error: %v", err)
		return
	}
	golog.Infof("GetSNSHttpEndpoint result: %s", result.UUID)
}

func TestUpdateSNSHttpEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSHttpEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSHttpEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSHttpEndpoint found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSHttpEndpointParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSHttpEndpointParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSHttpEndpoint(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSHttpEndpoint error: %v", err)
		return
	}
	golog.Infof("UpdateSNSHttpEndpoint result: %s", result.UUID)
}

func TestCreateSNSHttpEndpoint(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSHttpEndpoint is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSHttpEndpointParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSHttpEndpointParamDetail{
	// 		Name: "test-snshttpendpoint",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSHttpEndpoint(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSHttpEndpoint error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSHttpEndpoint result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSHttpEndpoint(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSHttpEndpoint error: %v", err)
	// }
}

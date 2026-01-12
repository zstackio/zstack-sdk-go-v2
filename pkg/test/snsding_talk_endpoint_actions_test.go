// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSDingTalkEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSDingTalkEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSDingTalkEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSDingTalkEndpoint result count: %d", len(result))
}

func TestUpdateSNSDingTalkEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSDingTalkEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSDingTalkEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSDingTalkEndpoint found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSDingTalkEndpointParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSDingTalkEndpointParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSDingTalkEndpoint(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSDingTalkEndpoint error: %v", err)
		return
	}
	golog.Infof("UpdateSNSDingTalkEndpoint result: %s", result.Uuid)
}

func TestCreateSNSDingTalkEndpoint(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSDingTalkEndpoint is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSDingTalkEndpointParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSDingTalkEndpointParamDetail{
	// 		Name: "test-snsdingtalkendpoint",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSDingTalkEndpoint(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSDingTalkEndpoint error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSDingTalkEndpoint result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSDingTalkEndpoint(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSDingTalkEndpoint error: %v", err)
	// }
}

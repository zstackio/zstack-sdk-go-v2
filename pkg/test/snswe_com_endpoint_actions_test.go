// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSWeComEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSWeComEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSWeComEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSWeComEndpoint result count: %d", len(result))
}

func TestUpdateSNSWeComEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSWeComEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSWeComEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSWeComEndpoint found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSWeComEndpointParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSWeComEndpointParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSWeComEndpoint(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSWeComEndpoint error: %v", err)
		return
	}
	golog.Infof("UpdateSNSWeComEndpoint result: %s", result.Uuid)
}

func TestCreateSNSWeComEndpoint(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSWeComEndpoint is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSWeComEndpointParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSWeComEndpointParamDetail{
	// 		Name: "test-snswecomendpoint",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSWeComEndpoint(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSWeComEndpoint error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSWeComEndpoint result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSWeComEndpoint(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSWeComEndpoint error: %v", err)
	// }
}

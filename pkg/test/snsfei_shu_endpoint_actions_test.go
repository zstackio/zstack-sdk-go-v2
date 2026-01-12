// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSFeiShuEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSFeiShuEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSFeiShuEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSFeiShuEndpoint result count: %d", len(result))
}

func TestUpdateSNSFeiShuEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSFeiShuEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSFeiShuEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSFeiShuEndpoint found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSFeiShuEndpointParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSFeiShuEndpointParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSFeiShuEndpoint(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSFeiShuEndpoint error: %v", err)
		return
	}
	golog.Infof("UpdateSNSFeiShuEndpoint result: %s", result.Uuid)
}

func TestCreateSNSFeiShuEndpoint(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSFeiShuEndpoint is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSFeiShuEndpointParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSFeiShuEndpointParamDetail{
	// 		Name: "test-snsfeishuendpoint",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSFeiShuEndpoint(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSFeiShuEndpoint error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSFeiShuEndpoint result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSFeiShuEndpoint(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSFeiShuEndpoint error: %v", err)
	// }
}

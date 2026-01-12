// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSUniversalSmsEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSUniversalSmsEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSUniversalSmsEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSUniversalSmsEndpoint result count: %d", len(result))
}

func TestUpdateSNSUniversalSmsEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSUniversalSmsEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSUniversalSmsEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSUniversalSmsEndpoint found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSUniversalSmsEndpointParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSUniversalSmsEndpointParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSUniversalSmsEndpoint(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSUniversalSmsEndpoint error: %v", err)
		return
	}
	golog.Infof("UpdateSNSUniversalSmsEndpoint result: %s", result.Uuid)
}

func TestCreateSNSUniversalSmsEndpoint(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSUniversalSmsEndpoint is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSUniversalSmsEndpointParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSUniversalSmsEndpointParamDetail{
	// 		Name: "test-snsuniversalsmsendpoint",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSUniversalSmsEndpoint(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSUniversalSmsEndpoint error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSUniversalSmsEndpoint result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSUniversalSmsEndpoint(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSUniversalSmsEndpoint error: %v", err)
	// }
}

func TestValidateSNSUniversalSmsEndpoint(t *testing.T) {
	// ValidateSNSUniversalSmsEndpoint operation
	t.Skip("TestValidateSNSUniversalSmsEndpoint requires manual implementation")

}

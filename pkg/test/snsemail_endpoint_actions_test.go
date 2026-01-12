// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSEmailEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSEmailEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSEmailEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSEmailEndpoint result count: %d", len(result))
}
func TestGetSNSEmailEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSEmailEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSEmailEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSEmailEndpoint found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSEmailEndpoint(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSEmailEndpoint error: %v", err)
		return
	}
	golog.Infof("GetSNSEmailEndpoint result: %s", result.UUID)
}

func TestCreateSNSEmailEndpoint(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSEmailEndpoint is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSEmailEndpointParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSEmailEndpointParamDetail{
	// 		Name: "test-snsemailendpoint",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSEmailEndpoint(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSEmailEndpoint error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSEmailEndpoint result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSEmailEndpoint(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSEmailEndpoint error: %v", err)
	// }
}

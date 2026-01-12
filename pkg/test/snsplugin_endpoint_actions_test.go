// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSPluginEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSPluginEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSPluginEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSPluginEndpoint result count: %d", len(result))
}

func TestCreateSNSPluginEndpoint(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSPluginEndpoint is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSPluginEndpointParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSPluginEndpointParamDetail{
	// 		Name: "test-snspluginendpoint",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSPluginEndpoint(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSPluginEndpoint error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSPluginEndpoint result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSPluginEndpoint(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSPluginEndpoint error: %v", err)
	// }
}

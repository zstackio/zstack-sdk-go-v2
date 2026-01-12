// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSMicrosoftTeamsEndpoint(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSMicrosoftTeamsEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSMicrosoftTeamsEndpoint error: %v", err)
		return
	}
	golog.Infof("QuerySNSMicrosoftTeamsEndpoint result count: %d", len(result))
}
func TestGetSNSMicrosoftTeamsEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSMicrosoftTeamsEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSMicrosoftTeamsEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSMicrosoftTeamsEndpoint found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSMicrosoftTeamsEndpoint(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSMicrosoftTeamsEndpoint error: %v", err)
		return
	}
	golog.Infof("GetSNSMicrosoftTeamsEndpoint result: %s", result.UUID)
}

func TestUpdateSNSMicrosoftTeamsEndpoint(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSMicrosoftTeamsEndpoint(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSNSMicrosoftTeamsEndpoint Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSMicrosoftTeamsEndpoint found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSNSMicrosoftTeamsEndpointParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSNSMicrosoftTeamsEndpointParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSNSMicrosoftTeamsEndpoint(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSNSMicrosoftTeamsEndpoint error: %v", err)
		return
	}
	golog.Infof("UpdateSNSMicrosoftTeamsEndpoint result: %s", result.UUID)
}

func TestCreateSNSMicrosoftTeamsEndpoint(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSNSMicrosoftTeamsEndpoint is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSNSMicrosoftTeamsEndpointParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSNSMicrosoftTeamsEndpointParamDetail{
	// 		Name: "test-snsmicrosoftteamsendpoint",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSNSMicrosoftTeamsEndpoint(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSNSMicrosoftTeamsEndpoint error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSNSMicrosoftTeamsEndpoint result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSNSMicrosoftTeamsEndpoint(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSNSMicrosoftTeamsEndpoint error: %v", err)
	// }
}

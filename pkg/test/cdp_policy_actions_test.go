// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCdpPolicy(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCdpPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCdpPolicy error: %v", err)
		return
	}
	golog.Infof("QueryCdpPolicy result count: %d", len(result))
}
func TestGetCdpPolicy(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCdpPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestGetCdpPolicy Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CdpPolicy found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetCdpPolicy(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCdpPolicy error: %v", err)
		return
	}
	golog.Infof("GetCdpPolicy result: %s", result.UUID)
}

func TestUpdateCdpPolicy(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCdpPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateCdpPolicy Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CdpPolicy found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateCdpPolicyParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateCdpPolicyParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateCdpPolicy(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateCdpPolicy error: %v", err)
		return
	}
	golog.Infof("UpdateCdpPolicy result: %s", result.UUID)
}

func TestDeleteCdpPolicy(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteCdpPolicy is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCdpPolicy(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteCdpPolicy Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CdpPolicy found to test Delete")
		return
	}

	err = accountLoginCli.DeleteCdpPolicy(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteCdpPolicy error: %v", err)
		return
	}
	golog.Infof("DeleteCdpPolicy succeeded for UUID: %s", list[0].UUID)
}

func TestCreateCdpPolicy(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateCdpPolicy is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateCdpPolicyParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateCdpPolicyParamDetail{
	// 		Name: "test-cdppolicy",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateCdpPolicy(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateCdpPolicy error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateCdpPolicy result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteCdpPolicy(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteCdpPolicy error: %v", err)
	// }
}

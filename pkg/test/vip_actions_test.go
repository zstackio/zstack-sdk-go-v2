// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVip(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVip(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVip error: %v", err)
		return
	}
	golog.Infof("QueryVip result count: %d", len(result))
}

func TestUpdateVip(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVip(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVip Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Vip found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVipParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVipParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVip(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVip error: %v", err)
		return
	}
	golog.Infof("UpdateVip result: %s", result.Uuid)
}

func TestDeleteVip(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVip is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVip(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVip Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Vip found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVip(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVip error: %v", err)
		return
	}
	golog.Infof("DeleteVip succeeded for UUID: %s", list[0].Uuid)
}

func TestCreateVip(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVip is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVipParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVipParamDetail{
	// 		Name: "test-vip",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVip(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVip error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVip result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVip(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVip error: %v", err)
	// }
}

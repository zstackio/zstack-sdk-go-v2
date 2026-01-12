// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVniRange(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVniRange(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVniRange error: %v", err)
		return
	}
	golog.Infof("QueryVniRange result count: %d", len(result))
}

func TestUpdateVniRange(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVniRange(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVniRange Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VniRange found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVniRangeParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVniRangeParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVniRange(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVniRange error: %v", err)
		return
	}
	golog.Infof("UpdateVniRange result: %s", result.Uuid)
}

func TestDeleteVniRange(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVniRange is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVniRange(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVniRange Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VniRange found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVniRange(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVniRange error: %v", err)
		return
	}
	golog.Infof("DeleteVniRange succeeded for UUID: %s", list[0].Uuid)
}

func TestCreateVniRange(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVniRange is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVniRangeParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVniRangeParamDetail{
	// 		Name: "test-vnirange",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVniRange(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVniRange error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVniRange result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVniRange(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVniRange error: %v", err)
	// }
}

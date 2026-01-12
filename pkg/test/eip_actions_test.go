// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEip(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryEip(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEip error: %v", err)
		return
	}
	golog.Infof("QueryEip result count: %d", len(result))
}
func TestGetEip(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEip(&queryParam)
	if err != nil {
		t.Errorf("TestGetEip Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Eip found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetEip(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetEip error: %v", err)
		return
	}
	golog.Infof("GetEip result: %s", result.UUID)
}

func TestUpdateEip(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEip(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateEip Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Eip found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateEipParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateEipParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateEip(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateEip error: %v", err)
		return
	}
	golog.Infof("UpdateEip result: %s", result.UUID)
}

func TestDeleteEip(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteEip is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryEip(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteEip Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Eip found to test Delete")
		return
	}

	err = accountLoginCli.DeleteEip(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteEip error: %v", err)
		return
	}
	golog.Infof("DeleteEip succeeded for UUID: %s", list[0].UUID)
}

func TestCreateEip(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateEip is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateEipParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateEipParamDetail{
	// 		Name: "test-eip",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateEip(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateEip error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateEip result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteEip(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteEip error: %v", err)
	// }
}

func TestAttachEip(t *testing.T) {
	// Attach operation
	t.Skip("TestAttachEip requires valid resource UUIDs to attach")

}

func TestDetachEip(t *testing.T) {
	// Detach operation
	t.Skip("TestDetachEip requires an attached resource")

}

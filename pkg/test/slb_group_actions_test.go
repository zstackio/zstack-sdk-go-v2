// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySlbGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySlbGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySlbGroup error: %v", err)
		return
	}
	golog.Infof("QuerySlbGroup result count: %d", len(result))
}

func TestUpdateSlbGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySlbGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSlbGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SlbGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSlbGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSlbGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSlbGroup(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSlbGroup error: %v", err)
		return
	}
	golog.Infof("UpdateSlbGroup result: %s", result.Uuid)
}

func TestDeleteSlbGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSlbGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySlbGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSlbGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SlbGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSlbGroup(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSlbGroup error: %v", err)
		return
	}
	golog.Infof("DeleteSlbGroup succeeded for UUID: %s", list[0].Uuid)
}

func TestCreateSlbGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateSlbGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateSlbGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateSlbGroupParamDetail{
	// 		Name: "test-slbgroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateSlbGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateSlbGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateSlbGroup result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteSlbGroup(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteSlbGroup error: %v", err)
	// }
}

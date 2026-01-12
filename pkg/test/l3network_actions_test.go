// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL3Network(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL3Network(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL3Network error: %v", err)
		return
	}
	golog.Infof("QueryL3Network result count: %d", len(result))
}
func TestGetL3Network(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL3Network(&queryParam)
	if err != nil {
		t.Errorf("TestGetL3Network Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L3Network found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetL3Network(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL3Network error: %v", err)
		return
	}
	golog.Infof("GetL3Network result: %s", result.UUID)
}

func TestUpdateL3Network(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL3Network(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateL3Network Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L3Network found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateL3NetworkParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateL3NetworkParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateL3Network(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateL3Network error: %v", err)
		return
	}
	golog.Infof("UpdateL3Network result: %s", result.UUID)
}

func TestDeleteL3Network(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteL3Network is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL3Network(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteL3Network Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L3Network found to test Delete")
		return
	}

	err = accountLoginCli.DeleteL3Network(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteL3Network error: %v", err)
		return
	}
	golog.Infof("DeleteL3Network succeeded for UUID: %s", list[0].UUID)
}

func TestCreateL3Network(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateL3Network is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateL3NetworkParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateL3NetworkParamDetail{
	// 		Name: "test-l3network",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateL3Network(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateL3Network error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateL3Network result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteL3Network(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteL3Network error: %v", err)
	// }
}

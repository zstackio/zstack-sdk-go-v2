// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDirectory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryDirectory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDirectory error: %v", err)
		return
	}
	golog.Infof("QueryDirectory result count: %d", len(result))
}
func TestGetDirectory(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDirectory(&queryParam)
	if err != nil {
		t.Errorf("TestGetDirectory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Directory found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetDirectory(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetDirectory error: %v", err)
		return
	}
	golog.Infof("GetDirectory result: %s", result.UUID)
}

func TestUpdateDirectory(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDirectory(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateDirectory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Directory found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateDirectoryParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateDirectoryParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateDirectory(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateDirectory error: %v", err)
		return
	}
	golog.Infof("UpdateDirectory result: %s", result.UUID)
}

func TestDeleteDirectory(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteDirectory is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDirectory(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteDirectory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Directory found to test Delete")
		return
	}

	err = accountLoginCli.DeleteDirectory(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteDirectory error: %v", err)
		return
	}
	golog.Infof("DeleteDirectory succeeded for UUID: %s", list[0].UUID)
}

func TestCreateDirectory(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateDirectory is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateDirectoryParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateDirectoryParamDetail{
	// 		Name: "test-directory",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateDirectory(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateDirectory error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateDirectory result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteDirectory(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteDirectory error: %v", err)
	// }
}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryResourceStack(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryResourceStack(&queryParam)
	if err != nil {
		t.Errorf("TestQueryResourceStack error: %v", err)
		return
	}
	golog.Infof("QueryResourceStack result count: %d", len(result))
}
func TestGetResourceStack(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryResourceStack(&queryParam)
	if err != nil {
		t.Errorf("TestGetResourceStack Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ResourceStack found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetResourceStack(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetResourceStack error: %v", err)
		return
	}
	golog.Infof("GetResourceStack result: %s", result.UUID)
}

func TestUpdateResourceStack(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryResourceStack(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateResourceStack Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ResourceStack found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateResourceStackParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateResourceStackParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateResourceStack(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateResourceStack error: %v", err)
		return
	}
	golog.Infof("UpdateResourceStack result: %s", result.UUID)
}

func TestDeleteResourceStack(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteResourceStack is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryResourceStack(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteResourceStack Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ResourceStack found to test Delete")
		return
	}

	err = accountLoginCli.DeleteResourceStack(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteResourceStack error: %v", err)
		return
	}
	golog.Infof("DeleteResourceStack succeeded for UUID: %s", list[0].UUID)
}

func TestCreateResourceStack(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateResourceStack is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateResourceStackParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateResourceStackParamDetail{
	// 		Name: "test-resourcestack",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateResourceStack(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateResourceStack error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateResourceStack result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteResourceStack(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteResourceStack error: %v", err)
	// }
}

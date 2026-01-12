// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPortMirror(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPortMirror(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPortMirror error: %v", err)
		return
	}
	golog.Infof("QueryPortMirror result count: %d", len(result))
}
func TestGetPortMirror(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPortMirror(&queryParam)
	if err != nil {
		t.Errorf("TestGetPortMirror Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PortMirror found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetPortMirror(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetPortMirror error: %v", err)
		return
	}
	golog.Infof("GetPortMirror result: %s", result.UUID)
}

func TestUpdatePortMirror(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPortMirror(&queryParam)
	if err != nil {
		t.Errorf("TestUpdatePortMirror Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PortMirror found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdatePortMirrorParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdatePortMirrorParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdatePortMirror(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdatePortMirror error: %v", err)
		return
	}
	golog.Infof("UpdatePortMirror result: %s", result.UUID)
}

func TestDeletePortMirror(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePortMirror is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPortMirror(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePortMirror Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PortMirror found to test Delete")
		return
	}

	err = accountLoginCli.DeletePortMirror(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePortMirror error: %v", err)
		return
	}
	golog.Infof("DeletePortMirror succeeded for UUID: %s", list[0].UUID)
}

func TestCreatePortMirror(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreatePortMirror is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreatePortMirrorParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreatePortMirrorParamDetail{
	// 		Name: "test-portmirror",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreatePortMirror(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreatePortMirror error: %v", err)
	// 	return
	// }
	// golog.Infof("CreatePortMirror result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeletePortMirror(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeletePortMirror error: %v", err)
	// }
}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2Network(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryL2Network(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2Network error: %v", err)
		return
	}
	golog.Infof("QueryL2Network result count: %d", len(result))
}
func TestGetL2Network(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL2Network(&queryParam)
	if err != nil {
		t.Errorf("TestGetL2Network Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2Network found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetL2Network(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetL2Network error: %v", err)
		return
	}
	golog.Infof("GetL2Network result: %s", result.UUID)
}

func TestUpdateL2Network(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL2Network(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateL2Network Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2Network found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateL2NetworkParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateL2NetworkParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateL2Network(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateL2Network error: %v", err)
		return
	}
	golog.Infof("UpdateL2Network result: %s", result.UUID)
}

func TestDeleteL2Network(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteL2Network is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryL2Network(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteL2Network Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No L2Network found to test Delete")
		return
	}

	err = accountLoginCli.DeleteL2Network(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteL2Network error: %v", err)
		return
	}
	golog.Infof("DeleteL2Network succeeded for UUID: %s", list[0].UUID)
}

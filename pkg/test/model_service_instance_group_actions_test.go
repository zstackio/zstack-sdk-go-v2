// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryModelServiceInstanceGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryModelServiceInstanceGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryModelServiceInstanceGroup error: %v", err)
		return
	}
	golog.Infof("QueryModelServiceInstanceGroup result count: %d", len(result))
}

func TestUpdateModelServiceInstanceGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelServiceInstanceGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateModelServiceInstanceGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelServiceInstanceGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateModelServiceInstanceGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateModelServiceInstanceGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateModelServiceInstanceGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateModelServiceInstanceGroup error: %v", err)
		return
	}
	golog.Infof("UpdateModelServiceInstanceGroup result: %s", result.UUID)
}

func TestDeleteModelServiceInstanceGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteModelServiceInstanceGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelServiceInstanceGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteModelServiceInstanceGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelServiceInstanceGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteModelServiceInstanceGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteModelServiceInstanceGroup error: %v", err)
		return
	}
	golog.Infof("DeleteModelServiceInstanceGroup succeeded for UUID: %s", list[0].UUID)
}

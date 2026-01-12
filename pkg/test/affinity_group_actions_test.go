// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAffinityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAffinityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAffinityGroup error: %v", err)
		return
	}
	golog.Infof("QueryAffinityGroup result count: %d", len(result))
}

func TestUpdateAffinityGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAffinityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAffinityGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AffinityGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAffinityGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAffinityGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAffinityGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAffinityGroup error: %v", err)
		return
	}
	golog.Infof("UpdateAffinityGroup result: %s", result.UUID)
}

func TestDeleteAffinityGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAffinityGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAffinityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAffinityGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AffinityGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAffinityGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAffinityGroup error: %v", err)
		return
	}
	golog.Infof("DeleteAffinityGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateAffinityGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateAffinityGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateAffinityGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateAffinityGroupParamDetail{
	// 		Name: "test-affinitygroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateAffinityGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateAffinityGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateAffinityGroup result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteAffinityGroup(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteAffinityGroup error: %v", err)
	// }
}

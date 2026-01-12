// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeSnapshotGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVolumeSnapshotGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeSnapshotGroup error: %v", err)
		return
	}
	golog.Infof("QueryVolumeSnapshotGroup result count: %d", len(result))
}
func TestGetVolumeSnapshotGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolumeSnapshotGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshotGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshotGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVolumeSnapshotGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshotGroup error: %v", err)
		return
	}
	golog.Infof("GetVolumeSnapshotGroup result: %s", result.UUID)
}

func TestUpdateVolumeSnapshotGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolumeSnapshotGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVolumeSnapshotGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshotGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVolumeSnapshotGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVolumeSnapshotGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVolumeSnapshotGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVolumeSnapshotGroup error: %v", err)
		return
	}
	golog.Infof("UpdateVolumeSnapshotGroup result: %s", result.UUID)
}

func TestDeleteVolumeSnapshotGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVolumeSnapshotGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolumeSnapshotGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVolumeSnapshotGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshotGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVolumeSnapshotGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVolumeSnapshotGroup error: %v", err)
		return
	}
	golog.Infof("DeleteVolumeSnapshotGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateVolumeSnapshotGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVolumeSnapshotGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVolumeSnapshotGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVolumeSnapshotGroupParamDetail{
	// 		Name: "test-volumesnapshotgroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVolumeSnapshotGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVolumeSnapshotGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVolumeSnapshotGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVolumeSnapshotGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVolumeSnapshotGroup error: %v", err)
	// }
}

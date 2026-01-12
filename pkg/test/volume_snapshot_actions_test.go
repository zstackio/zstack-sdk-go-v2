// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeSnapshot(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVolumeSnapshot(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeSnapshot error: %v", err)
		return
	}
	golog.Infof("QueryVolumeSnapshot result count: %d", len(result))
}
func TestGetVolumeSnapshot(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolumeSnapshot(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshot Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshot found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVolumeSnapshot(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolumeSnapshot error: %v", err)
		return
	}
	golog.Infof("GetVolumeSnapshot result: %s", result.UUID)
}

func TestUpdateVolumeSnapshot(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolumeSnapshot(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateVolumeSnapshot Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshot found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateVolumeSnapshotParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateVolumeSnapshotParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateVolumeSnapshot(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateVolumeSnapshot error: %v", err)
		return
	}
	golog.Infof("UpdateVolumeSnapshot result: %s", result.UUID)
}

func TestDeleteVolumeSnapshot(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVolumeSnapshot is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolumeSnapshot(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVolumeSnapshot Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeSnapshot found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVolumeSnapshot(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVolumeSnapshot error: %v", err)
		return
	}
	golog.Infof("DeleteVolumeSnapshot succeeded for UUID: %s", list[0].UUID)
}

func TestCreateVolumeSnapshot(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVolumeSnapshot is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVolumeSnapshotParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVolumeSnapshotParamDetail{
	// 		Name: "test-volumesnapshot",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVolumeSnapshot(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVolumeSnapshot error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVolumeSnapshot result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVolumeSnapshot(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVolumeSnapshot error: %v", err)
	// }
}

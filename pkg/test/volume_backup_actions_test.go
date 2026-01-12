// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolumeBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVolumeBackup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolumeBackup error: %v", err)
		return
	}
	golog.Infof("QueryVolumeBackup result count: %d", len(result))
}
func TestGetVolumeBackup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolumeBackup(&queryParam)
	if err != nil {
		t.Errorf("TestGetVolumeBackup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeBackup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVolumeBackup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVolumeBackup error: %v", err)
		return
	}
	golog.Infof("GetVolumeBackup result: %s", result.UUID)
}

func TestDeleteVolumeBackup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVolumeBackup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVolumeBackup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVolumeBackup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VolumeBackup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVolumeBackup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVolumeBackup error: %v", err)
		return
	}
	golog.Infof("DeleteVolumeBackup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateVolumeBackup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateVolumeBackup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateVolumeBackupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateVolumeBackupParamDetail{
	// 		Name: "test-volumebackup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateVolumeBackup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateVolumeBackup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateVolumeBackup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteVolumeBackup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteVolumeBackup error: %v", err)
	// }
}

func TestSyncVolumeBackup(t *testing.T) {
	// Sync operation
	t.Skip("TestSyncVolumeBackup requires a valid resource to sync")

}

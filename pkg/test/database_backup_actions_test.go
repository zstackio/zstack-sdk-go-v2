// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDatabaseBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryDatabaseBackup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDatabaseBackup error: %v", err)
		return
	}
	golog.Infof("QueryDatabaseBackup result count: %d", len(result))
}

func TestDeleteDatabaseBackup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteDatabaseBackup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDatabaseBackup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteDatabaseBackup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No DatabaseBackup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteDatabaseBackup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteDatabaseBackup error: %v", err)
		return
	}
	golog.Infof("DeleteDatabaseBackup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateDatabaseBackup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateDatabaseBackup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateDatabaseBackupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateDatabaseBackupParamDetail{
	// 		Name: "test-databasebackup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateDatabaseBackup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateDatabaseBackup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateDatabaseBackup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteDatabaseBackup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteDatabaseBackup error: %v", err)
	// }
}

func TestSyncDatabaseBackup(t *testing.T) {
	// Sync operation
	t.Skip("TestSyncDatabaseBackup requires a valid resource to sync")

}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryBackupStorage result count: %d", len(result))
}

func TestUpdateBackupStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBackupStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BackupStorage found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBackupStorageParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBackupStorageParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBackupStorage(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBackupStorage error: %v", err)
		return
	}
	golog.Infof("UpdateBackupStorage result: %s", result.UUID)
}

func TestDeleteBackupStorage(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteBackupStorage is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteBackupStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BackupStorage found to test Delete")
		return
	}

	err = accountLoginCli.DeleteBackupStorage(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteBackupStorage error: %v", err)
		return
	}
	golog.Infof("DeleteBackupStorage succeeded for UUID: %s", list[0].UUID)
}

func TestReconnectBackupStorage(t *testing.T) {
	// ReconnectBackupStorage operation
	t.Skip("TestReconnectBackupStorage requires manual implementation")

}

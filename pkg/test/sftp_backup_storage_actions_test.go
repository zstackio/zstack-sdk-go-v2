// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySftpBackupStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySftpBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySftpBackupStorage error: %v", err)
		return
	}
	golog.Infof("QuerySftpBackupStorage result count: %d", len(result))
}
func TestGetSftpBackupStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySftpBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetSftpBackupStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SftpBackupStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSftpBackupStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSftpBackupStorage error: %v", err)
		return
	}
	golog.Infof("GetSftpBackupStorage result: %s", result.UUID)
}

func TestUpdateSftpBackupStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySftpBackupStorage(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSftpBackupStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SftpBackupStorage found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSftpBackupStorageParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSftpBackupStorageParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSftpBackupStorage(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSftpBackupStorage error: %v", err)
		return
	}
	golog.Infof("UpdateSftpBackupStorage result: %s", result.UUID)
}

func TestReconnectSftpBackupStorage(t *testing.T) {
	// ReconnectSftpBackupStorage operation
	t.Skip("TestReconnectSftpBackupStorage requires manual implementation")

}

func TestAddSftpBackupStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSftpBackupStorage requires valid creation parameters")

}

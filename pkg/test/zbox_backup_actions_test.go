// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryZBoxBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryZBoxBackup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryZBoxBackup error: %v", err)
		return
	}
	golog.Infof("QueryZBoxBackup result count: %d", len(result))
}
func TestGetZBoxBackup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryZBoxBackup(&queryParam)
	if err != nil {
		t.Errorf("TestGetZBoxBackup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ZBoxBackup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetZBoxBackup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetZBoxBackup error: %v", err)
		return
	}
	golog.Infof("GetZBoxBackup result: %s", result.UUID)
}

func TestCreateZBoxBackup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateZBoxBackup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateZBoxBackupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateZBoxBackupParamDetail{
	// 		Name: "test-zboxbackup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateZBoxBackup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateZBoxBackup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateZBoxBackup result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteZBoxBackup(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteZBoxBackup error: %v", err)
	// }
}

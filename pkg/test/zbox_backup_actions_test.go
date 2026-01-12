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

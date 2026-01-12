// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNfvInstGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNfvInstGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNfvInstGroup error: %v", err)
		return
	}
	golog.Infof("QueryNfvInstGroup result count: %d", len(result))
}

func TestUpdateNfvInstGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNfvInstGroup(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateNfvInstGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NfvInstGroup found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateNfvInstGroupParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateNfvInstGroupParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateNfvInstGroup(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateNfvInstGroup error: %v", err)
		return
	}
	golog.Infof("UpdateNfvInstGroup result: %s", result.UUID)
}

func TestDeleteNfvInstGroup(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteNfvInstGroup is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNfvInstGroup(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteNfvInstGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NfvInstGroup found to test Delete")
		return
	}

	err = accountLoginCli.DeleteNfvInstGroup(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteNfvInstGroup error: %v", err)
		return
	}
	golog.Infof("DeleteNfvInstGroup succeeded for UUID: %s", list[0].UUID)
}

func TestCreateNfvInstGroup(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateNfvInstGroup is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateNfvInstGroupParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateNfvInstGroupParamDetail{
	// 		Name: "test-nfvinstgroup",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateNfvInstGroup(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateNfvInstGroup error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateNfvInstGroup result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteNfvInstGroup(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteNfvInstGroup error: %v", err)
	// }
}

func TestSyncNfvInstGroup(t *testing.T) {
	// Sync operation
	t.Skip("TestSyncNfvInstGroup requires a valid resource to sync")

}

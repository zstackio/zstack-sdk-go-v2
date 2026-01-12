// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryPrimaryStorage result count: %d", len(result))
}

func TestUpdatePrimaryStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestUpdatePrimaryStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PrimaryStorage found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdatePrimaryStorageParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdatePrimaryStorageParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdatePrimaryStorage(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdatePrimaryStorage error: %v", err)
		return
	}
	golog.Infof("UpdatePrimaryStorage result: %s", result.UUID)
}

func TestDeletePrimaryStorage(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeletePrimaryStorage is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestDeletePrimaryStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No PrimaryStorage found to test Delete")
		return
	}

	err = accountLoginCli.DeletePrimaryStorage(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeletePrimaryStorage error: %v", err)
		return
	}
	golog.Infof("DeletePrimaryStorage succeeded for UUID: %s", list[0].UUID)
}

func TestReconnectPrimaryStorage(t *testing.T) {
	// ReconnectPrimaryStorage operation
	t.Skip("TestReconnectPrimaryStorage requires manual implementation")

}

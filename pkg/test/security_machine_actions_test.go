// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySecurityMachine(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySecurityMachine(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySecurityMachine error: %v", err)
		return
	}
	golog.Infof("QuerySecurityMachine result count: %d", len(result))
}
func TestGetSecurityMachine(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecurityMachine(&queryParam)
	if err != nil {
		t.Errorf("TestGetSecurityMachine Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecurityMachine found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSecurityMachine(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSecurityMachine error: %v", err)
		return
	}
	golog.Infof("GetSecurityMachine result: %s", result.UUID)
}

func TestUpdateSecurityMachine(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecurityMachine(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSecurityMachine Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecurityMachine found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSecurityMachineParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSecurityMachineParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSecurityMachine(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSecurityMachine error: %v", err)
		return
	}
	golog.Infof("UpdateSecurityMachine result: %s", result.UUID)
}

func TestDeleteSecurityMachine(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSecurityMachine is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecurityMachine(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSecurityMachine Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecurityMachine found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSecurityMachine(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSecurityMachine error: %v", err)
		return
	}
	golog.Infof("DeleteSecurityMachine succeeded for UUID: %s", list[0].UUID)
}

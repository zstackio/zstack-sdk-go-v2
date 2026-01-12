// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLdapServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLdapServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLdapServer error: %v", err)
		return
	}
	golog.Infof("QueryLdapServer result count: %d", len(result))
}
func TestGetLdapServer(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLdapServer(&queryParam)
	if err != nil {
		t.Errorf("TestGetLdapServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LdapServer found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetLdapServer(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLdapServer error: %v", err)
		return
	}
	golog.Infof("GetLdapServer result: %s", result.UUID)
}

func TestUpdateLdapServer(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLdapServer(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateLdapServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LdapServer found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateLdapServerParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateLdapServerParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateLdapServer(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateLdapServer error: %v", err)
		return
	}
	golog.Infof("UpdateLdapServer result: %s", result.UUID)
}

func TestDeleteLdapServer(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteLdapServer is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLdapServer(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteLdapServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LdapServer found to test Delete")
		return
	}

	err = accountLoginCli.DeleteLdapServer(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteLdapServer error: %v", err)
		return
	}
	golog.Infof("DeleteLdapServer succeeded for UUID: %s", list[0].UUID)
}

func TestAddLdapServer(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddLdapServer requires valid creation parameters")

}

func TestSyncLdapServer(t *testing.T) {
	// Sync operation
	t.Skip("TestSyncLdapServer requires a valid resource to sync")

}

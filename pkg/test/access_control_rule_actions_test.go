// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccessControlRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAccessControlRule(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccessControlRule error: %v", err)
		return
	}
	golog.Infof("QueryAccessControlRule result count: %d", len(result))
}

func TestUpdateAccessControlRule(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccessControlRule(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateAccessControlRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccessControlRule found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateAccessControlRuleParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateAccessControlRuleParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateAccessControlRule(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateAccessControlRule error: %v", err)
		return
	}
	golog.Infof("UpdateAccessControlRule result: %s", result.UUID)
}

func TestDeleteAccessControlRule(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAccessControlRule is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccessControlRule(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAccessControlRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccessControlRule found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAccessControlRule(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAccessControlRule error: %v", err)
		return
	}
	golog.Infof("DeleteAccessControlRule succeeded for UUID: %s", list[0].UUID)
}

func TestAddAccessControlRule(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddAccessControlRule requires valid creation parameters")

}

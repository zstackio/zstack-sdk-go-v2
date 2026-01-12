// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNasMountTarget(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNasMountTarget(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNasMountTarget error: %v", err)
		return
	}
	golog.Infof("QueryNasMountTarget result count: %d", len(result))
}
func TestGetNasMountTarget(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNasMountTarget(&queryParam)
	if err != nil {
		t.Errorf("TestGetNasMountTarget Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NasMountTarget found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNasMountTarget(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNasMountTarget error: %v", err)
		return
	}
	golog.Infof("GetNasMountTarget result: %s", result.UUID)
}

func TestUpdateNasMountTarget(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNasMountTarget(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateNasMountTarget Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NasMountTarget found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateNasMountTargetParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateNasMountTargetParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateNasMountTarget(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateNasMountTarget error: %v", err)
		return
	}
	golog.Infof("UpdateNasMountTarget result: %s", result.UUID)
}

func TestDeleteNasMountTarget(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteNasMountTarget is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNasMountTarget(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteNasMountTarget Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NasMountTarget found to test Delete")
		return
	}

	err = accountLoginCli.DeleteNasMountTarget(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteNasMountTarget error: %v", err)
		return
	}
	golog.Infof("DeleteNasMountTarget succeeded for UUID: %s", list[0].UUID)
}

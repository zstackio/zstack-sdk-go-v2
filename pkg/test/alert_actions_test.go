// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAlert(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAlert(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAlert error: %v", err)
		return
	}
	golog.Infof("QueryAlert result count: %d", len(result))
}
func TestGetAlert(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAlert(&queryParam)
	if err != nil {
		t.Errorf("TestGetAlert Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Alert found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAlert(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAlert error: %v", err)
		return
	}
	golog.Infof("GetAlert result: %s", result.UUID)
}

func TestDeleteAlert(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteAlert is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAlert(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAlert Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Alert found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAlert(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAlert error: %v", err)
		return
	}
	golog.Infof("DeleteAlert succeeded for UUID: %s", list[0].UUID)
}

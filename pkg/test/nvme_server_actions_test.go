// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNvmeServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNvmeServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNvmeServer error: %v", err)
		return
	}
	golog.Infof("QueryNvmeServer result count: %d", len(result))
}
func TestGetNvmeServer(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNvmeServer(&queryParam)
	if err != nil {
		t.Errorf("TestGetNvmeServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NvmeServer found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNvmeServer(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNvmeServer error: %v", err)
		return
	}
	golog.Infof("GetNvmeServer result: %s", result.UUID)
}

func TestDeleteNvmeServer(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteNvmeServer is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNvmeServer(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteNvmeServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NvmeServer found to test Delete")
		return
	}

	err = accountLoginCli.DeleteNvmeServer(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteNvmeServer error: %v", err)
		return
	}
	golog.Infof("DeleteNvmeServer succeeded for UUID: %s", list[0].UUID)
}

func TestAddNvmeServer(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddNvmeServer requires valid creation parameters")

}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNvmeTarget(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNvmeTarget(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNvmeTarget error: %v", err)
		return
	}
	golog.Infof("QueryNvmeTarget result count: %d", len(result))
}
func TestGetNvmeTarget(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNvmeTarget(&queryParam)
	if err != nil {
		t.Errorf("TestGetNvmeTarget Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NvmeTarget found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNvmeTarget(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNvmeTarget error: %v", err)
		return
	}
	golog.Infof("GetNvmeTarget result: %s", result.UUID)
}

func TestRefreshNvmeTarget(t *testing.T) {
	// RefreshNvmeTarget operation
	t.Skip("TestRefreshNvmeTarget requires manual implementation")

}

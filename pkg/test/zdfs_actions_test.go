// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryZdfs(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryZdfs(&queryParam)
	if err != nil {
		t.Errorf("TestQueryZdfs error: %v", err)
		return
	}
	golog.Infof("QueryZdfs result count: %d", len(result))
}
func TestGetZdfs(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryZdfs(&queryParam)
	if err != nil {
		t.Errorf("TestGetZdfs Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Zdfs found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetZdfs(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetZdfs error: %v", err)
		return
	}
	golog.Infof("GetZdfs result: %s", result.UUID)
}

func TestReconnectZdfs(t *testing.T) {
	// ReconnectZdfs operation
	t.Skip("TestReconnectZdfs requires manual implementation")

}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephOsdGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCephOsdGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephOsdGroup error: %v", err)
		return
	}
	golog.Infof("QueryCephOsdGroup result count: %d", len(result))
}
func TestGetCephOsdGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCephOsdGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetCephOsdGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephOsdGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetCephOsdGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCephOsdGroup error: %v", err)
		return
	}
	golog.Infof("GetCephOsdGroup result: %s", result.UUID)
}

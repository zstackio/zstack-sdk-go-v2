// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostOsCategory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHostOsCategory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostOsCategory error: %v", err)
		return
	}
	golog.Infof("QueryHostOsCategory result count: %d", len(result))
}
func TestGetHostOsCategory(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHostOsCategory(&queryParam)
	if err != nil {
		t.Errorf("TestGetHostOsCategory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HostOsCategory found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetHostOsCategory(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetHostOsCategory error: %v", err)
		return
	}
	golog.Infof("GetHostOsCategory result: %s", result.UUID)
}

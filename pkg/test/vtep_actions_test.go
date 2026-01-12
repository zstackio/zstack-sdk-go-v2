// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVtep(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVtep(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVtep error: %v", err)
		return
	}
	golog.Infof("QueryVtep result count: %d", len(result))
}
func TestGetVtep(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVtep(&queryParam)
	if err != nil {
		t.Errorf("TestGetVtep Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Vtep found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVtep(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVtep error: %v", err)
		return
	}
	golog.Infof("GetVtep result: %s", result.UUID)
}

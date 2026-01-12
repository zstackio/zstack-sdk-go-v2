// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMttyDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMttyDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMttyDevice error: %v", err)
		return
	}
	golog.Infof("QueryMttyDevice result count: %d", len(result))
}
func TestGetMttyDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMttyDevice(&queryParam)
	if err != nil {
		t.Errorf("TestGetMttyDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MttyDevice found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMttyDevice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMttyDevice error: %v", err)
		return
	}
	golog.Infof("GetMttyDevice result: %s", result.UUID)
}

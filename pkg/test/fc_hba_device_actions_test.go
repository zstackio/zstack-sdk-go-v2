// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryFcHbaDevice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryFcHbaDevice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryFcHbaDevice error: %v", err)
		return
	}
	golog.Infof("QueryFcHbaDevice result count: %d", len(result))
}
func TestGetFcHbaDevice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryFcHbaDevice(&queryParam)
	if err != nil {
		t.Errorf("TestGetFcHbaDevice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No FcHbaDevice found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetFcHbaDevice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetFcHbaDevice error: %v", err)
		return
	}
	golog.Infof("GetFcHbaDevice result: %s", result.UUID)
}

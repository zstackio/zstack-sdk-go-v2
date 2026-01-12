// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryFiberChannelLun(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryFiberChannelLun(&queryParam)
	if err != nil {
		t.Errorf("TestQueryFiberChannelLun error: %v", err)
		return
	}
	golog.Infof("QueryFiberChannelLun result count: %d", len(result))
}
func TestGetFiberChannelLun(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryFiberChannelLun(&queryParam)
	if err != nil {
		t.Errorf("TestGetFiberChannelLun Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No FiberChannelLun found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetFiberChannelLun(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetFiberChannelLun error: %v", err)
		return
	}
	golog.Infof("GetFiberChannelLun result: %s", result.UUID)
}

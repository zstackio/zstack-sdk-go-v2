// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryFiberChannelStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryFiberChannelStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryFiberChannelStorage error: %v", err)
		return
	}
	golog.Infof("QueryFiberChannelStorage result count: %d", len(result))
}
func TestGetFiberChannelStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryFiberChannelStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetFiberChannelStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No FiberChannelStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetFiberChannelStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetFiberChannelStorage error: %v", err)
		return
	}
	golog.Infof("GetFiberChannelStorage result: %s", result.UUID)
}

func TestRefreshFiberChannelStorage(t *testing.T) {
	// RefreshFiberChannelStorage operation
	t.Skip("TestRefreshFiberChannelStorage requires manual implementation")

}

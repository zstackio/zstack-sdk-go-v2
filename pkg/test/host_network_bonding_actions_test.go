// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostNetworkBonding(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHostNetworkBonding(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostNetworkBonding error: %v", err)
		return
	}
	golog.Infof("QueryHostNetworkBonding result count: %d", len(result))
}
func TestGetHostNetworkBonding(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryHostNetworkBonding(&queryParam)
	if err != nil {
		t.Errorf("TestGetHostNetworkBonding Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No HostNetworkBonding found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetHostNetworkBonding(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetHostNetworkBonding error: %v", err)
		return
	}
	golog.Infof("GetHostNetworkBonding result: %s", result.UUID)
}

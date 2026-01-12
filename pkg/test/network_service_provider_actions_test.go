// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNetworkServiceProvider(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNetworkServiceProvider(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNetworkServiceProvider error: %v", err)
		return
	}
	golog.Infof("QueryNetworkServiceProvider result count: %d", len(result))
}
func TestGetNetworkServiceProvider(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNetworkServiceProvider(&queryParam)
	if err != nil {
		t.Errorf("TestGetNetworkServiceProvider Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NetworkServiceProvider found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNetworkServiceProvider(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNetworkServiceProvider error: %v", err)
		return
	}
	golog.Infof("GetNetworkServiceProvider result: %s", result.UUID)
}

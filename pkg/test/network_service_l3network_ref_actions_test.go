// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNetworkServiceL3NetworkRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNetworkServiceL3NetworkRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNetworkServiceL3NetworkRef error: %v", err)
		return
	}
	golog.Infof("QueryNetworkServiceL3NetworkRef result count: %d", len(result))
}
func TestGetNetworkServiceL3NetworkRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNetworkServiceL3NetworkRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetNetworkServiceL3NetworkRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NetworkServiceL3NetworkRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNetworkServiceL3NetworkRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNetworkServiceL3NetworkRef error: %v", err)
		return
	}
	golog.Infof("GetNetworkServiceL3NetworkRef result: %s", result.UUID)
}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcHaGroupNetworkServiceRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVpcHaGroupNetworkServiceRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcHaGroupNetworkServiceRef error: %v", err)
		return
	}
	golog.Infof("QueryVpcHaGroupNetworkServiceRef result count: %d", len(result))
}
func TestGetVpcHaGroupNetworkServiceRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcHaGroupNetworkServiceRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetVpcHaGroupNetworkServiceRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcHaGroupNetworkServiceRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVpcHaGroupNetworkServiceRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVpcHaGroupNetworkServiceRef error: %v", err)
		return
	}
	golog.Infof("GetVpcHaGroupNetworkServiceRef result: %s", result.UUID)
}

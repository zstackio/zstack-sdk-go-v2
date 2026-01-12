// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcSnatState(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVpcSnatState(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcSnatState error: %v", err)
		return
	}
	golog.Infof("QueryVpcSnatState result count: %d", len(result))
}
func TestGetVpcSnatState(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcSnatState(&queryParam)
	if err != nil {
		t.Errorf("TestGetVpcSnatState Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcSnatState found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVpcSnatState(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVpcSnatState error: %v", err)
		return
	}
	golog.Infof("GetVpcSnatState result: %s", result.UUID)
}

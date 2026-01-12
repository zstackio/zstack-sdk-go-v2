// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterResourcePool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenterResourcePool(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterResourcePool error: %v", err)
		return
	}
	golog.Infof("QueryVCenterResourcePool result count: %d", len(result))
}
func TestGetVCenterResourcePool(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVCenterResourcePool(&queryParam)
	if err != nil {
		t.Errorf("TestGetVCenterResourcePool Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenterResourcePool found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVCenterResourcePool(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVCenterResourcePool error: %v", err)
		return
	}
	golog.Infof("GetVCenterResourcePool result: %s", result.UUID)
}

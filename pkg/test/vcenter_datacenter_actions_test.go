// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterDatacenter(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenterDatacenter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterDatacenter error: %v", err)
		return
	}
	golog.Infof("QueryVCenterDatacenter result count: %d", len(result))
}
func TestGetVCenterDatacenter(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVCenterDatacenter(&queryParam)
	if err != nil {
		t.Errorf("TestGetVCenterDatacenter Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenterDatacenter found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVCenterDatacenter(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVCenterDatacenter error: %v", err)
		return
	}
	golog.Infof("GetVCenterDatacenter result: %s", result.UUID)
}

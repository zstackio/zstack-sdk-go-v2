// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterCluster(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenterCluster(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterCluster error: %v", err)
		return
	}
	golog.Infof("QueryVCenterCluster result count: %d", len(result))
}
func TestGetVCenterCluster(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVCenterCluster(&queryParam)
	if err != nil {
		t.Errorf("TestGetVCenterCluster Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VCenterCluster found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVCenterCluster(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVCenterCluster error: %v", err)
		return
	}
	golog.Infof("GetVCenterCluster result: %s", result.UUID)
}

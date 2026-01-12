// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNativeCluster(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNativeCluster(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNativeCluster error: %v", err)
		return
	}
	golog.Infof("QueryNativeCluster result count: %d", len(result))
}
func TestGetNativeCluster(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNativeCluster(&queryParam)
	if err != nil {
		t.Errorf("TestGetNativeCluster Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NativeCluster found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNativeCluster(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNativeCluster error: %v", err)
		return
	}
	golog.Infof("GetNativeCluster result: %s", result.UUID)
}

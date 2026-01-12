// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNativeHost(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNativeHost(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNativeHost error: %v", err)
		return
	}
	golog.Infof("QueryNativeHost result count: %d", len(result))
}
func TestGetNativeHost(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryNativeHost(&queryParam)
	if err != nil {
		t.Errorf("TestGetNativeHost Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No NativeHost found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetNativeHost(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetNativeHost error: %v", err)
		return
	}
	golog.Infof("GetNativeHost result: %s", result.UUID)
}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImageCache(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImageCache(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImageCache error: %v", err)
		return
	}
	golog.Infof("QueryImageCache result count: %d", len(result))
}
func TestGetImageCache(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImageCache(&queryParam)
	if err != nil {
		t.Errorf("TestGetImageCache Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ImageCache found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetImageCache(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetImageCache error: %v", err)
		return
	}
	golog.Infof("GetImageCache result: %s", result.UUID)
}

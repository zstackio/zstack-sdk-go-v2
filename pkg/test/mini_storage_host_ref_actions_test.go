// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMiniStorageHostRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMiniStorageHostRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMiniStorageHostRef error: %v", err)
		return
	}
	golog.Infof("QueryMiniStorageHostRef result count: %d", len(result))
}
func TestGetMiniStorageHostRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMiniStorageHostRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetMiniStorageHostRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MiniStorageHostRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMiniStorageHostRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMiniStorageHostRef error: %v", err)
		return
	}
	golog.Infof("GetMiniStorageHostRef result: %s", result.UUID)
}

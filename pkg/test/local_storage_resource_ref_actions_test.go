// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLocalStorageResourceRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLocalStorageResourceRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLocalStorageResourceRef error: %v", err)
		return
	}
	golog.Infof("QueryLocalStorageResourceRef result count: %d", len(result))
}
func TestGetLocalStorageResourceRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLocalStorageResourceRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetLocalStorageResourceRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LocalStorageResourceRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetLocalStorageResourceRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLocalStorageResourceRef error: %v", err)
		return
	}
	golog.Infof("GetLocalStorageResourceRef result: %s", result.UUID)
}

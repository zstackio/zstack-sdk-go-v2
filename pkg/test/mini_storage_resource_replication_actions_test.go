// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMiniStorageResourceReplication(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMiniStorageResourceReplication(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMiniStorageResourceReplication error: %v", err)
		return
	}
	golog.Infof("QueryMiniStorageResourceReplication result count: %d", len(result))
}
func TestGetMiniStorageResourceReplication(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMiniStorageResourceReplication(&queryParam)
	if err != nil {
		t.Errorf("TestGetMiniStorageResourceReplication Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MiniStorageResourceReplication found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMiniStorageResourceReplication(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMiniStorageResourceReplication error: %v", err)
		return
	}
	golog.Infof("GetMiniStorageResourceReplication result: %s", result.UUID)
}

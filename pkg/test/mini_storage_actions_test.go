// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMiniStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMiniStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMiniStorage error: %v", err)
		return
	}
	golog.Infof("QueryMiniStorage result count: %d", len(result))
}
func TestGetMiniStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMiniStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetMiniStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No MiniStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMiniStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMiniStorage error: %v", err)
		return
	}
	golog.Infof("GetMiniStorage result: %s", result.UUID)
}

func TestAddMiniStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddMiniStorage requires valid creation parameters")

}

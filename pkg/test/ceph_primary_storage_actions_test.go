// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCephPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCephPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCephPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryCephPrimaryStorage result count: %d", len(result))
}
func TestGetCephPrimaryStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCephPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetCephPrimaryStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CephPrimaryStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetCephPrimaryStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCephPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("GetCephPrimaryStorage result: %s", result.UUID)
}

func TestAddCephPrimaryStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddCephPrimaryStorage requires valid creation parameters")

}

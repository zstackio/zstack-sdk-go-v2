// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImageGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImageGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImageGroup error: %v", err)
		return
	}
	golog.Infof("QueryImageGroup result count: %d", len(result))
}
func TestGetImageGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImageGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetImageGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ImageGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetImageGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetImageGroup error: %v", err)
		return
	}
	golog.Infof("GetImageGroup result: %s", result.UUID)
}

func TestExpungeImageGroup(t *testing.T) {
	// Expunge operation - permanently deletes
	t.Skip("TestExpungeImageGroup is dangerous - permanently deletes resource")

}

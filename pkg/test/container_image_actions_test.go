// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryContainerImage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryContainerImage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryContainerImage error: %v", err)
		return
	}
	golog.Infof("QueryContainerImage result count: %d", len(result))
}
func TestGetContainerImage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryContainerImage(&queryParam)
	if err != nil {
		t.Errorf("TestGetContainerImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ContainerImage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetContainerImage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetContainerImage error: %v", err)
		return
	}
	golog.Infof("GetContainerImage result: %s", result.UUID)
}

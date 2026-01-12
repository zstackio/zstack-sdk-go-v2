// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryKvmHypervisorInfo(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryKvmHypervisorInfo(&queryParam)
	if err != nil {
		t.Errorf("TestQueryKvmHypervisorInfo error: %v", err)
		return
	}
	golog.Infof("QueryKvmHypervisorInfo result count: %d", len(result))
}
func TestGetKvmHypervisorInfo(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryKvmHypervisorInfo(&queryParam)
	if err != nil {
		t.Errorf("TestGetKvmHypervisorInfo Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No KvmHypervisorInfo found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetKvmHypervisorInfo(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetKvmHypervisorInfo error: %v", err)
		return
	}
	golog.Infof("GetKvmHypervisorInfo result: %s", result.UUID)
}

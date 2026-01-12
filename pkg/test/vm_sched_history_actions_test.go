// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmSchedHistory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmSchedHistory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmSchedHistory error: %v", err)
		return
	}
	golog.Infof("QueryVmSchedHistory result count: %d", len(result))
}
func TestGetVmSchedHistory(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVmSchedHistory(&queryParam)
	if err != nil {
		t.Errorf("TestGetVmSchedHistory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VmSchedHistory found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVmSchedHistory(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVmSchedHistory error: %v", err)
		return
	}
	golog.Infof("GetVmSchedHistory result: %s", result.UUID)
}

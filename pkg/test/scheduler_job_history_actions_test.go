// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySchedulerJobHistory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySchedulerJobHistory(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySchedulerJobHistory error: %v", err)
		return
	}
	golog.Infof("QuerySchedulerJobHistory result count: %d", len(result))
}
func TestGetSchedulerJobHistory(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySchedulerJobHistory(&queryParam)
	if err != nil {
		t.Errorf("TestGetSchedulerJobHistory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SchedulerJobHistory found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSchedulerJobHistory(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSchedulerJobHistory error: %v", err)
		return
	}
	golog.Infof("GetSchedulerJobHistory result: %s", result.UUID)
}

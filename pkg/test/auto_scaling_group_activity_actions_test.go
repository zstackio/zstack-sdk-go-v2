// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAutoScalingGroupActivity(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAutoScalingGroupActivity(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAutoScalingGroupActivity error: %v", err)
		return
	}
	golog.Infof("QueryAutoScalingGroupActivity result count: %d", len(result))
}
func TestGetAutoScalingGroupActivity(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAutoScalingGroupActivity(&queryParam)
	if err != nil {
		t.Errorf("TestGetAutoScalingGroupActivity Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AutoScalingGroupActivity found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAutoScalingGroupActivity(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAutoScalingGroupActivity error: %v", err)
		return
	}
	golog.Infof("GetAutoScalingGroupActivity result: %s", result.UUID)
}

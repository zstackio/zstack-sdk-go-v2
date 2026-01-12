// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryModelEvalServiceInstanceGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryModelEvalServiceInstanceGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryModelEvalServiceInstanceGroup error: %v", err)
		return
	}
	golog.Infof("QueryModelEvalServiceInstanceGroup result count: %d", len(result))
}
func TestGetModelEvalServiceInstanceGroup(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryModelEvalServiceInstanceGroup(&queryParam)
	if err != nil {
		t.Errorf("TestGetModelEvalServiceInstanceGroup Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ModelEvalServiceInstanceGroup found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetModelEvalServiceInstanceGroup(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetModelEvalServiceInstanceGroup error: %v", err)
		return
	}
	golog.Infof("GetModelEvalServiceInstanceGroup result: %s", result.UUID)
}

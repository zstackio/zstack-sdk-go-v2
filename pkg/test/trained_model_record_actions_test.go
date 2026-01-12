// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTrainedModelRecord(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryTrainedModelRecord(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTrainedModelRecord error: %v", err)
		return
	}
	golog.Infof("QueryTrainedModelRecord result count: %d", len(result))
}
func TestGetTrainedModelRecord(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryTrainedModelRecord(&queryParam)
	if err != nil {
		t.Errorf("TestGetTrainedModelRecord Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No TrainedModelRecord found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetTrainedModelRecord(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetTrainedModelRecord error: %v", err)
		return
	}
	golog.Infof("GetTrainedModelRecord result: %s", result.UUID)
}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryExponBlockVolume(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryExponBlockVolume(&queryParam)
	if err != nil {
		t.Errorf("TestQueryExponBlockVolume error: %v", err)
		return
	}
	golog.Infof("QueryExponBlockVolume result count: %d", len(result))
}
func TestGetExponBlockVolume(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryExponBlockVolume(&queryParam)
	if err != nil {
		t.Errorf("TestGetExponBlockVolume Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ExponBlockVolume found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetExponBlockVolume(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetExponBlockVolume error: %v", err)
		return
	}
	golog.Infof("GetExponBlockVolume result: %s", result.UUID)
}

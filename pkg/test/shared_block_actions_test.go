// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedBlock(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySharedBlock(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedBlock error: %v", err)
		return
	}
	golog.Infof("QuerySharedBlock result count: %d", len(result))
}

func TestUpdateSharedBlock(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySharedBlock(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateSharedBlock Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SharedBlock found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateSharedBlockParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateSharedBlockParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateSharedBlock("", list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateSharedBlock error: %v", err)
		return
	}
	golog.Infof("UpdateSharedBlock result: %s", result.UUID)
}

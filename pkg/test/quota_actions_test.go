// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryQuota(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryQuota(&queryParam)
	if err != nil {
		t.Errorf("TestQueryQuota error: %v", err)
		return
	}
	golog.Infof("QueryQuota result count: %d", len(result))
}

func TestUpdateQuota(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryQuota(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateQuota Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Quota found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateQuotaParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateQuotaParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateQuota(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateQuota error: %v", err)
		return
	}
	golog.Infof("UpdateQuota result: %s", result.UUID)
}

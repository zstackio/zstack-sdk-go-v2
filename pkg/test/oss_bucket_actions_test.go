// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateOssBucket(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryOssBucketFileName(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateOssBucket Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No OssBucket found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateOssBucketParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateOssBucketParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateOssBucket(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateOssBucket error: %v", err)
		return
	}
	golog.Infof("UpdateOssBucket result: %s", result.UUID)
}

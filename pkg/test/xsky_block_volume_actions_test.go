// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryXskyBlockVolume(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryXskyBlockVolume(&queryParam)
	if err != nil {
		t.Errorf("TestQueryXskyBlockVolume error: %v", err)
		return
	}
	golog.Infof("QueryXskyBlockVolume result count: %d", len(result))
}

func TestUpdateXskyBlockVolume(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryXskyBlockVolume(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateXskyBlockVolume Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No XskyBlockVolume found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateXskyBlockVolumeParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateXskyBlockVolumeParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateXskyBlockVolume(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateXskyBlockVolume error: %v", err)
		return
	}
	golog.Infof("UpdateXskyBlockVolume result: %s", result.Uuid)
}

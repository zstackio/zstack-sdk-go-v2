// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGlobalConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGlobalConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGlobalConfig error: %v", err)
		return
	}
	golog.Infof("QueryGlobalConfig result count: %d", len(result))
}

func TestUpdateGlobalConfig(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryGlobalConfig(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateGlobalConfig Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No GlobalConfig found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateGlobalConfigParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateGlobalConfigParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateGlobalConfig(list[0].Name, updateParam)
	if err != nil {
		t.Errorf("TestUpdateGlobalConfig error: %v", err)
		return
	}
	golog.Infof("UpdateGlobalConfig result: %s", result.Name)
}

func TestResetGlobalConfig(t *testing.T) {
	// Reset operation
	t.Skip("TestResetGlobalConfig may affect resource state")

}

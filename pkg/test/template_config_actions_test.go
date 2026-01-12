// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryTemplateConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryTemplateConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryTemplateConfig error: %v", err)
		return
	}
	golog.Infof("QueryTemplateConfig result count: %d", len(result))
}

func TestUpdateTemplateConfig(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryTemplateConfig(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateTemplateConfig Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No TemplateConfig found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateTemplateConfigParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateTemplateConfigParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateTemplateConfig(*list[0].TemplateUuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateTemplateConfig error: %v", err)
		return
	}
	golog.Infof("UpdateTemplateConfig result: %s", result.TemplateUuid)
}

func TestRevertTemplateConfig(t *testing.T) {
	// RevertTemplateConfig operation
	t.Skip("TestRevertTemplateConfig requires manual implementation")

}

func TestApplyTemplateConfig(t *testing.T) {
	// ApplyTemplateConfig operation
	t.Skip("TestApplyTemplateConfig requires manual implementation")

}

func TestResetTemplateConfig(t *testing.T) {
	// Reset operation
	t.Skip("TestResetTemplateConfig may affect resource state")

}

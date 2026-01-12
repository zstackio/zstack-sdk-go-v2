// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2ProjectAttribute(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2ProjectAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2ProjectAttribute error: %v", err)
		return
	}
	golog.Infof("QueryIAM2ProjectAttribute result count: %d", len(result))
}

func TestUpdateIAM2ProjectAttribute(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2ProjectAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2ProjectAttribute Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2ProjectAttribute found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIAM2ProjectAttributeParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIAM2ProjectAttributeParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIAM2ProjectAttribute(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2ProjectAttribute error: %v", err)
		return
	}
	golog.Infof("UpdateIAM2ProjectAttribute result: %s", result.UUID)
}

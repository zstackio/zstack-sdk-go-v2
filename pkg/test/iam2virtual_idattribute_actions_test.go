// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2VirtualIDAttribute(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2VirtualIDAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2VirtualIDAttribute error: %v", err)
		return
	}
	golog.Infof("QueryIAM2VirtualIDAttribute result count: %d", len(result))
}

func TestUpdateIAM2VirtualIDAttribute(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2VirtualIDAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2VirtualIDAttribute Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualIDAttribute found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIAM2VirtualIDAttributeParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIAM2VirtualIDAttributeParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIAM2VirtualIDAttribute(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2VirtualIDAttribute error: %v", err)
		return
	}
	golog.Infof("UpdateIAM2VirtualIDAttribute result: %s", result.UUID)
}

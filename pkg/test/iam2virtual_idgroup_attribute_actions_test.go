// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2VirtualIDGroupAttribute(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2VirtualIDGroupAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2VirtualIDGroupAttribute error: %v", err)
		return
	}
	golog.Infof("QueryIAM2VirtualIDGroupAttribute result count: %d", len(result))
}
func TestGetIAM2VirtualIDGroupAttribute(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2VirtualIDGroupAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2VirtualIDGroupAttribute Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualIDGroupAttribute found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetIAM2VirtualIDGroupAttribute(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2VirtualIDGroupAttribute error: %v", err)
		return
	}
	golog.Infof("GetIAM2VirtualIDGroupAttribute result: %s", result.UUID)
}

func TestUpdateIAM2VirtualIDGroupAttribute(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2VirtualIDGroupAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2VirtualIDGroupAttribute Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2VirtualIDGroupAttribute found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIAM2VirtualIDGroupAttributeParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIAM2VirtualIDGroupAttributeParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIAM2VirtualIDGroupAttribute(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2VirtualIDGroupAttribute error: %v", err)
		return
	}
	golog.Infof("UpdateIAM2VirtualIDGroupAttribute result: %s", result.UUID)
}

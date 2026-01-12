// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2OrganizationAttribute(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2OrganizationAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2OrganizationAttribute error: %v", err)
		return
	}
	golog.Infof("QueryIAM2OrganizationAttribute result count: %d", len(result))
}
func TestGetIAM2OrganizationAttribute(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2OrganizationAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2OrganizationAttribute Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2OrganizationAttribute found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetIAM2OrganizationAttribute(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2OrganizationAttribute error: %v", err)
		return
	}
	golog.Infof("GetIAM2OrganizationAttribute result: %s", result.UUID)
}

func TestUpdateIAM2OrganizationAttribute(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2OrganizationAttribute(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2OrganizationAttribute Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2OrganizationAttribute found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIAM2OrganizationAttributeParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIAM2OrganizationAttributeParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIAM2OrganizationAttribute(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIAM2OrganizationAttribute error: %v", err)
		return
	}
	golog.Infof("UpdateIAM2OrganizationAttribute result: %s", result.UUID)
}

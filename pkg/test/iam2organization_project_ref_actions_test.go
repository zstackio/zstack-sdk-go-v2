// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2OrganizationProjectRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2OrganizationProjectRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2OrganizationProjectRef error: %v", err)
		return
	}
	golog.Infof("QueryIAM2OrganizationProjectRef result count: %d", len(result))
}
func TestGetIAM2OrganizationProjectRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2OrganizationProjectRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2OrganizationProjectRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2OrganizationProjectRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetIAM2OrganizationProjectRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2OrganizationProjectRef error: %v", err)
		return
	}
	golog.Infof("GetIAM2OrganizationProjectRef result: %s", result.UUID)
}

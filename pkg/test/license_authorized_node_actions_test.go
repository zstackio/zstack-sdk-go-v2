// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLicenseAuthorizedNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryLicenseAuthorizedNode(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLicenseAuthorizedNode error: %v", err)
		return
	}
	golog.Infof("QueryLicenseAuthorizedNode result count: %d", len(result))
}
func TestGetLicenseAuthorizedNode(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLicenseAuthorizedNode(&queryParam)
	if err != nil {
		t.Errorf("TestGetLicenseAuthorizedNode Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LicenseAuthorizedNode found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetLicenseAuthorizedNode(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLicenseAuthorizedNode error: %v", err)
		return
	}
	golog.Infof("GetLicenseAuthorizedNode result: %s", result.UUID)
}

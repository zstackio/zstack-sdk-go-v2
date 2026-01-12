// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestGetLicenseAuthorizedCapacity(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLicenseAuthorizedNode(&queryParam)
	if err != nil {
		t.Errorf("TestGetLicenseAuthorizedCapacity Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LicenseAuthorizedCapacity found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetLicenseAuthorizedCapacity(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLicenseAuthorizedCapacity error: %v", err)
		return
	}
	golog.Infof("GetLicenseAuthorizedCapacity result: %s", result.Id)
}

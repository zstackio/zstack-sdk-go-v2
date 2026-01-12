// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestUpdateLicense(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLicenseAuthorizedNode(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateLicense Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No License found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateLicenseParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateLicenseParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateLicense(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateLicense error: %v", err)
		return
	}
	golog.Infof("UpdateLicense result: %s", result.UUID)
}

func TestDeleteLicense(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteLicense is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryLicenseAuthorizedNode(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteLicense Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No License found to test Delete")
		return
	}

	err = accountLoginCli.DeleteLicense(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteLicense error: %v", err)
		return
	}
	golog.Infof("DeleteLicense succeeded for UUID: %s", list[0].UUID)
}

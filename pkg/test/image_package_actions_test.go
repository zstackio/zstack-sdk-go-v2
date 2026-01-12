// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImagePackage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImagePackage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImagePackage error: %v", err)
		return
	}
	golog.Infof("QueryImagePackage result count: %d", len(result))
}

func TestUpdateImagePackage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImagePackage(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateImagePackage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ImagePackage found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateImagePackageParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateImagePackageParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateImagePackage(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateImagePackage error: %v", err)
		return
	}
	golog.Infof("UpdateImagePackage result: %s", result.UUID)
}

func TestDeleteImagePackage(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteImagePackage is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImagePackage(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteImagePackage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ImagePackage found to test Delete")
		return
	}

	err = accountLoginCli.DeleteImagePackage(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteImagePackage error: %v", err)
		return
	}
	golog.Infof("DeleteImagePackage succeeded for UUID: %s", list[0].UUID)
}

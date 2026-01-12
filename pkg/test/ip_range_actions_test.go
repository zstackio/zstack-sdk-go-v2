// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIpRange(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIpRange(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIpRange error: %v", err)
		return
	}
	golog.Infof("QueryIpRange result count: %d", len(result))
}

func TestUpdateIpRange(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIpRange(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateIpRange Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IpRange found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateIpRangeParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateIpRangeParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateIpRange(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateIpRange error: %v", err)
		return
	}
	golog.Infof("UpdateIpRange result: %s", result.UUID)
}

func TestDeleteIpRange(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteIpRange is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIpRange(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteIpRange Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IpRange found to test Delete")
		return
	}

	err = accountLoginCli.DeleteIpRange(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteIpRange error: %v", err)
		return
	}
	golog.Infof("DeleteIpRange succeeded for UUID: %s", list[0].UUID)
}

func TestAddIpRange(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddIpRange requires valid creation parameters")

}

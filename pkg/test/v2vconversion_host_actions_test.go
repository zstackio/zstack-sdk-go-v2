// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryV2VConversionHost(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryV2VConversionHost(&queryParam)
	if err != nil {
		t.Errorf("TestQueryV2VConversionHost error: %v", err)
		return
	}
	golog.Infof("QueryV2VConversionHost result count: %d", len(result))
}

func TestUpdateV2VConversionHost(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryV2VConversionHost(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateV2VConversionHost Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No V2VConversionHost found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateV2VConversionHostParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateV2VConversionHostParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateV2VConversionHost(list[0].Uuid, updateParam)
	if err != nil {
		t.Errorf("TestUpdateV2VConversionHost error: %v", err)
		return
	}
	golog.Infof("UpdateV2VConversionHost result: %s", result.Uuid)
}

func TestDeleteV2VConversionHost(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteV2VConversionHost is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryV2VConversionHost(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteV2VConversionHost Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No V2VConversionHost found to test Delete")
		return
	}

	err = accountLoginCli.DeleteV2VConversionHost(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteV2VConversionHost error: %v", err)
		return
	}
	golog.Infof("DeleteV2VConversionHost succeeded for UUID: %s", list[0].Uuid)
}

func TestAddV2VConversionHost(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddV2VConversionHost requires valid creation parameters")

}

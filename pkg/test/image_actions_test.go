// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImage error: %v", err)
		return
	}
	golog.Infof("QueryImage result count: %d", len(result))
}

func TestGetImage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestGetImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetImage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetImage error: %v", err)
		return
	}
	golog.Infof("GetImage result: %s", result.UUID)
}

func TestUpdateImage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateImageParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateImageParamDetail{
			Name: "centos-test",
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateImage(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateImage error: %v", err)
		return
	}
	golog.Infof("UpdateImage result: %s", result.UUID)
}

func TestDeleteImage(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteImage is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImage(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteImage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Image found to test Delete")
		return
	}

	err = accountLoginCli.DeleteImage(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteImage error: %v", err)
		return
	}
	golog.Infof("DeleteImage succeeded for UUID: %s", list[0].UUID)
}

func TestAddImage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddImage requires valid creation parameters")

}

func TestAddImageAsync(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddImageAsync requires valid creation parameters")

}

func TestSyncImage(t *testing.T) {
	// Sync operation
	t.Skip("TestSyncImage requires a valid resource to sync")

}

func TestRecoverImage(t *testing.T) {
	// Recover operation - requires a deleted resource
	t.Skip("TestRecoverImage requires a deleted resource UUID")

}

func TestCloneImage(t *testing.T) {
	// Clone operation
	t.Skip("TestCloneImage requires a valid resource to clone")

}

func TestExpungeImage(t *testing.T) {
	// Expunge operation - permanently deletes
	t.Skip("TestExpungeImage is dangerous - permanently deletes resource")

}

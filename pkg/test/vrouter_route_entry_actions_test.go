// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVRouterRouteEntry(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVRouterRouteEntry(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVRouterRouteEntry error: %v", err)
		return
	}
	golog.Infof("QueryVRouterRouteEntry result count: %d", len(result))
}

func TestDeleteVRouterRouteEntry(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteVRouterRouteEntry is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVRouterRouteEntry(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteVRouterRouteEntry Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VRouterRouteEntry found to test Delete")
		return
	}

	err = accountLoginCli.DeleteVRouterRouteEntry(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteVRouterRouteEntry error: %v", err)
		return
	}
	golog.Infof("DeleteVRouterRouteEntry succeeded for UUID: %s", list[0].Uuid)
}

func TestAddVRouterRouteEntry(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddVRouterRouteEntry requires valid creation parameters")

}

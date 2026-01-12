// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryShareableVolumeVmInstanceRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryShareableVolumeVmInstanceRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryShareableVolumeVmInstanceRef error: %v", err)
		return
	}
	golog.Infof("QueryShareableVolumeVmInstanceRef result count: %d", len(result))
}
func TestGetShareableVolumeVmInstanceRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryShareableVolumeVmInstanceRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetShareableVolumeVmInstanceRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ShareableVolumeVmInstanceRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetShareableVolumeVmInstanceRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetShareableVolumeVmInstanceRef error: %v", err)
		return
	}
	golog.Infof("GetShareableVolumeVmInstanceRef result: %s", result.UUID)
}

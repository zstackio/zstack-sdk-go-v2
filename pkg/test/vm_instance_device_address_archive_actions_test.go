// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstanceDeviceAddressArchive(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmInstanceDeviceAddressArchive(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstanceDeviceAddressArchive error: %v", err)
		return
	}
	golog.Infof("QueryVmInstanceDeviceAddressArchive result count: %d", len(result))
}

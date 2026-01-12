// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstancePciDeviceSpecRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmInstancePciDeviceSpecRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstancePciDeviceSpecRef error: %v", err)
		return
	}
	golog.Infof("QueryVmInstancePciDeviceSpecRef result count: %d", len(result))
}

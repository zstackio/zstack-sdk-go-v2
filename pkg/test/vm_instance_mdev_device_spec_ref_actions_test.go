// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVmInstanceMdevDeviceSpecRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVmInstanceMdevDeviceSpecRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstanceMdevDeviceSpecRef error: %v", err)
		return
	}
	golog.Infof("QueryVmInstanceMdevDeviceSpecRef result count: %d", len(result))
}


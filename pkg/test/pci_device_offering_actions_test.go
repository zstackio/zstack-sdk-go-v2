// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPciDeviceOffering(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPciDeviceOffering(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPciDeviceOffering error: %v", err)
		return
	}
	golog.Infof("QueryPciDeviceOffering result count: %d", len(result))
}


// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryScsiLun(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryScsiLun(&queryParam)
	if err != nil {
		t.Errorf("TestQueryScsiLun error: %v", err)
		return
	}
	golog.Infof("QueryScsiLun result count: %d", len(result))
}


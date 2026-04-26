// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroupTemplateRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorGroupTemplateRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroupTemplateRef error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroupTemplateRef result count: %d", len(result))
}


// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMonitorGroupInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMonitorGroupInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMonitorGroupInstance error: %v", err)
		return
	}
	golog.Infof("QueryMonitorGroupInstance result count: %d", len(result))
}

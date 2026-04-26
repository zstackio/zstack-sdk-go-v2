// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPhysicalSwitch(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPhysicalSwitch(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPhysicalSwitch error: %v", err)
		return
	}
	golog.Infof("QueryPhysicalSwitch result count: %d", len(result))
}


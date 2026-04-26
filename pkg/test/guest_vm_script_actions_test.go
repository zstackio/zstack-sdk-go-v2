// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGuestVmScript(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGuestVmScript(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGuestVmScript error: %v", err)
		return
	}
	golog.Infof("QueryGuestVmScript result count: %d", len(result))
}


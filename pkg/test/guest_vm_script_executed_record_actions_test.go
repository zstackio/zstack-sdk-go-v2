// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGuestVmScriptExecutedRecord(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGuestVmScriptExecutedRecord(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGuestVmScriptExecutedRecord error: %v", err)
		return
	}
	golog.Infof("QueryGuestVmScriptExecutedRecord result count: %d", len(result))
}

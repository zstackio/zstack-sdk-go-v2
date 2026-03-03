// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryGuestVmScriptExecutedRecordDetail(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryGuestVmScriptExecutedRecordDetail(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGuestVmScriptExecutedRecordDetail error: %v", err)
		return
	}
	golog.Infof("QueryGuestVmScriptExecutedRecordDetail result count: %d", len(result))
}


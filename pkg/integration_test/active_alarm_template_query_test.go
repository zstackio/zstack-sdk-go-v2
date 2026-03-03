// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryActiveAlarmTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryActiveAlarmTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryActiveAlarmTemplate error: %v", err)
		return
	}
	golog.Infof("QueryActiveAlarmTemplate result count: %d", len(result))
}


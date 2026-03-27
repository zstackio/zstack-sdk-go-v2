// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPortMirror(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryPortMirror(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPortMirror error: %v", err)
		return
	}
	golog.Infof("QueryPortMirror result count: %d", len(result))
}


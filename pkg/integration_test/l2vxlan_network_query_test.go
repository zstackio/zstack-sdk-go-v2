// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryL2VxlanNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryL2VxlanNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VxlanNetwork error: %v", err)
		return
	}
	golog.Infof("QueryL2VxlanNetwork result count: %d", len(result))
}


// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySshKeyPair(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySshKeyPair(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySshKeyPair error: %v", err)
		return
	}
	golog.Infof("QuerySshKeyPair result count: %d", len(result))
}


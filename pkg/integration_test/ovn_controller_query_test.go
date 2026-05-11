// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryOvnController(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryOvnController(&queryParam)
	if err != nil {
		t.Errorf("TestQueryOvnController error: %v", err)
		return
	}
	golog.Infof("QueryOvnController result count: %d", len(result))
}


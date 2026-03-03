// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHuaweiIMasterFabric(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryHuaweiIMasterFabric(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHuaweiIMasterFabric error: %v", err)
		return
	}
	golog.Infof("QueryHuaweiIMasterFabric result count: %d", len(result))
}


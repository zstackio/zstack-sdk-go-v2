// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySecretResourcePool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySecretResourcePool(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySecretResourcePool error: %v", err)
		return
	}
	golog.Infof("QuerySecretResourcePool result count: %d", len(result))
}


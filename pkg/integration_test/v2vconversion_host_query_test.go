// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryV2VConversionHost(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryV2VConversionHost(&queryParam)
	if err != nil {
		t.Errorf("TestQueryV2VConversionHost error: %v", err)
		return
	}
	golog.Infof("QueryV2VConversionHost result count: %d", len(result))
}


// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVolume(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVolume(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVolume error: %v", err)
		return
	}
	golog.Infof("QueryVolume result count: %d", len(result))
}


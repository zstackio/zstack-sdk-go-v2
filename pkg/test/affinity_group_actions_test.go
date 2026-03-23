// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAffinityGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAffinityGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAffinityGroup error: %v", err)
		return
	}
	golog.Infof("QueryAffinityGroup result count: %d", len(result))
}


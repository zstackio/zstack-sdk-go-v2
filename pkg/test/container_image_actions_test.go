// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryContainerImage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryContainerImage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryContainerImage error: %v", err)
		return
	}
	golog.Infof("QueryContainerImage result count: %d", len(result))
}


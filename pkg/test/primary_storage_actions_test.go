// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryPrimaryStorage result count: %d", len(result))
}


// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSTopic(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSTopic(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSTopic error: %v", err)
		return
	}
	golog.Infof("QuerySNSTopic result count: %d", len(result))
}


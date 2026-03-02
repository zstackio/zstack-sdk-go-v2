// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDRSAdvice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryDRSAdvice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDRSAdvice error: %v", err)
		return
	}
	golog.Infof("QueryDRSAdvice result count: %d", len(result))
}


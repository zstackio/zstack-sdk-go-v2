// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPriceTable(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPriceTable(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPriceTable error: %v", err)
		return
	}
	golog.Infof("QueryPriceTable result count: %d", len(result))
}


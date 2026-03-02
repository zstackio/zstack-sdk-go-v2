// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccountPriceTableRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAccountPriceTableRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccountPriceTableRef error: %v", err)
		return
	}
	golog.Infof("QueryAccountPriceTableRef result count: %d", len(result))
}


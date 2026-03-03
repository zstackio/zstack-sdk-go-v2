// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryUser(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryUser(&queryParam)
	if err != nil {
		t.Errorf("TestQueryUser error: %v", err)
		return
	}
	golog.Infof("QueryUser result count: %d", len(result))
}


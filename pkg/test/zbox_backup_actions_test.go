// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryZBoxBackup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryZBoxBackup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryZBoxBackup error: %v", err)
		return
	}
	golog.Infof("QueryZBoxBackup result count: %d", len(result))
}


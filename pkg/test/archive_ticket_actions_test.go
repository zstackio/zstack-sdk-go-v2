// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryArchiveTicket(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryArchiveTicket(&queryParam)
	if err != nil {
		t.Errorf("TestQueryArchiveTicket error: %v", err)
		return
	}
	golog.Infof("QueryArchiveTicket result count: %d", len(result))
}


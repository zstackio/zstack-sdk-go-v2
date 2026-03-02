// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2Organization(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2Organization(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2Organization error: %v", err)
		return
	}
	golog.Infof("QueryIAM2Organization result count: %d", len(result))
}


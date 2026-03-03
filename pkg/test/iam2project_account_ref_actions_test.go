// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2ProjectAccountRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2ProjectAccountRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2ProjectAccountRef error: %v", err)
		return
	}
	golog.Infof("QueryIAM2ProjectAccountRef result count: %d", len(result))
}


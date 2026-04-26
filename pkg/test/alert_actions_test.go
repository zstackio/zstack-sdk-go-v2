// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAlert(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAlert(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAlert error: %v", err)
		return
	}
	golog.Infof("QueryAlert result count: %d", len(result))
}


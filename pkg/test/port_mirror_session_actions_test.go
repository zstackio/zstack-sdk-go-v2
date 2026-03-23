// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryPortMirrorSession(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryPortMirrorSession(&queryParam)
	if err != nil {
		t.Errorf("TestQueryPortMirrorSession error: %v", err)
		return
	}
	golog.Infof("QueryPortMirrorSession result count: %d", len(result))
}


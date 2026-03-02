// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcSharedQos(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVpcSharedQos(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcSharedQos error: %v", err)
		return
	}
	golog.Infof("QueryVpcSharedQos result count: %d", len(result))
}


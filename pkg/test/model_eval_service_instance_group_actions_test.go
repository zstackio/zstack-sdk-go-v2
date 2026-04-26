// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryModelEvalServiceInstanceGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryModelEvalServiceInstanceGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryModelEvalServiceInstanceGroup error: %v", err)
		return
	}
	golog.Infof("QueryModelEvalServiceInstanceGroup result count: %d", len(result))
}


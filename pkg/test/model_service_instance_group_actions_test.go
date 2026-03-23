// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryModelServiceInstanceGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryModelServiceInstanceGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryModelServiceInstanceGroup error: %v", err)
		return
	}
	golog.Infof("QueryModelServiceInstanceGroup result count: %d", len(result))
}


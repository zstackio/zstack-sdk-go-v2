// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryApplianceVm(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryApplianceVm(&queryParam)
	if err != nil {
		t.Errorf("TestQueryApplianceVm error: %v", err)
		return
	}
	golog.Infof("QueryApplianceVm result count: %d", len(result))
}


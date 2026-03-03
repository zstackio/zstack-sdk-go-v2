// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHuaweiIMasterVpc(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHuaweiIMasterVpc(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHuaweiIMasterVpc error: %v", err)
		return
	}
	golog.Infof("QueryHuaweiIMasterVpc result count: %d", len(result))
}


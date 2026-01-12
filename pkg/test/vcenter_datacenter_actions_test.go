// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterDatacenter(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenterDatacenter(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterDatacenter error: %v", err)
		return
	}
	golog.Infof("QueryVCenterDatacenter result count: %d", len(result))
}

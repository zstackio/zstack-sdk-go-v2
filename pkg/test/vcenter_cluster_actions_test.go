// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterCluster(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenterCluster(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterCluster error: %v", err)
		return
	}
	golog.Infof("QueryVCenterCluster result count: %d", len(result))
}

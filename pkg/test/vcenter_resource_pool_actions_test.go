// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVCenterResourcePool(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVCenterResourcePool(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterResourcePool error: %v", err)
		return
	}
	golog.Infof("QueryVCenterResourcePool result count: %d", len(result))
}

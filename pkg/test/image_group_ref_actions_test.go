// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImageGroupRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImageGroupRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImageGroupRef error: %v", err)
		return
	}
	golog.Infof("QueryImageGroupRef result count: %d", len(result))
}

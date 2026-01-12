// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySlbVmInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySlbVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySlbVmInstance error: %v", err)
		return
	}
	golog.Infof("QuerySlbVmInstance result count: %d", len(result))
}

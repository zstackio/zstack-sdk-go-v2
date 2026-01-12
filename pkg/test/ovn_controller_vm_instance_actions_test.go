// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryOvnControllerVmInstance(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryOvnControllerVmInstance(&queryParam)
	if err != nil {
		t.Errorf("TestQueryOvnControllerVmInstance error: %v", err)
		return
	}
	golog.Infof("QueryOvnControllerVmInstance result count: %d", len(result))
}

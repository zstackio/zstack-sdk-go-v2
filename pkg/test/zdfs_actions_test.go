// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryZdfs(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryZdfs(&queryParam)
	if err != nil {
		t.Errorf("TestQueryZdfs error: %v", err)
		return
	}
	golog.Infof("QueryZdfs result count: %d", len(result))
}

func TestReconnectZdfs(t *testing.T) {
	// ReconnectZdfs operation
	t.Skip("TestReconnectZdfs requires manual implementation")

}

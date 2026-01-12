// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNvmeTarget(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNvmeTarget(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNvmeTarget error: %v", err)
		return
	}
	golog.Infof("QueryNvmeTarget result count: %d", len(result))
}

func TestRefreshNvmeTarget(t *testing.T) {
	// RefreshNvmeTarget operation
	t.Skip("TestRefreshNvmeTarget requires manual implementation")

}

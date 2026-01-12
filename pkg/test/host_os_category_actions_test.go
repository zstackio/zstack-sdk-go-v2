// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryHostOsCategory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryHostOsCategory(&queryParam)
	if err != nil {
		t.Errorf("TestQueryHostOsCategory error: %v", err)
		return
	}
	golog.Infof("QueryHostOsCategory result count: %d", len(result))
}

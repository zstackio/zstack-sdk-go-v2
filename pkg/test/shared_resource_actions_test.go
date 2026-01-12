// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedResource(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySharedResource(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedResource error: %v", err)
		return
	}
	golog.Infof("QuerySharedResource result count: %d", len(result))
}

// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryNetworkServiceL3NetworkRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryNetworkServiceL3NetworkRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryNetworkServiceL3NetworkRef error: %v", err)
		return
	}
	golog.Infof("QueryNetworkServiceL3NetworkRef result count: %d", len(result))
}

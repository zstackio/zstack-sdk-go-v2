// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcHaGroupNetworkServiceRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVpcHaGroupNetworkServiceRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcHaGroupNetworkServiceRef error: %v", err)
		return
	}
	golog.Infof("QueryVpcHaGroupNetworkServiceRef result count: %d", len(result))
}

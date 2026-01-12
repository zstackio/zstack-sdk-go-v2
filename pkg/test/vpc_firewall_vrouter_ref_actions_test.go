// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryVpcFirewallVRouterRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryVpcFirewallVRouterRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcFirewallVRouterRef error: %v", err)
		return
	}
	golog.Infof("QueryVpcFirewallVRouterRef result count: %d", len(result))
}

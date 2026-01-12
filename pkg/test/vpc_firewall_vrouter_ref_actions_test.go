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
func TestGetVpcFirewallVRouterRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryVpcFirewallVRouterRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetVpcFirewallVRouterRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No VpcFirewallVRouterRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetVpcFirewallVRouterRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetVpcFirewallVRouterRef error: %v", err)
		return
	}
	golog.Infof("GetVpcFirewallVRouterRef result: %s", result.UUID)
}

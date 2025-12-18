// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLoadBalancerListenerACLEntries 获取LoadBalancerListenerACLEntries详情
func (cli *ZSClient) GetLoadBalancerListenerACLEntries(uuid string) (*view.GetLoadBalancerListenerACLEntriesView, error) {
	var resp view.GetLoadBalancerListenerACLEntriesView
	if err := cli.Get("v1/load-balancers/listeners/access-control-lists/entries", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}


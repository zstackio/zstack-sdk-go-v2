// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLoadBalancerListenerACLEntries gets LoadBalancerListenerACLEntries by uuid
func (cli *ZSClient) GetLoadBalancerListenerACLEntries(uuid string) (*view.GetLoadBalancerListenerACLEntriesView, error) {
	var resp view.GetLoadBalancerListenerACLEntriesView
	if err := cli.Get("v1/load-balancers/listeners/access-control-lists/entries", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

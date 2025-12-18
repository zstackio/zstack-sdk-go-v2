// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddServerGroupToLoadBalancerListener 操作AddServerGroupToLoadBalancerListener
func (cli *ZSClient) AddServerGroupToLoadBalancerListener(params param.AddServerGroupToLoadBalancerListenerParam) (*view.AddServerGroupToLoadBalancerListenerEventView, error) {
	resp := view.AddServerGroupToLoadBalancerListenerEventView{}
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/servergroups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}


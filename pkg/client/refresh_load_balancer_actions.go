// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshLoadBalancer 操作RefreshLoadBalancer
func (cli *ZSClient) RefreshLoadBalancer(uuid string, params param.RefreshLoadBalancerParam) (*view.RefreshLoadBalancerEventView, error) {
	resp := view.RefreshLoadBalancerEventView{}
	if err := cli.Put("v1/load-balancers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}


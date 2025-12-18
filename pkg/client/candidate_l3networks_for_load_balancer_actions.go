// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateL3NetworksForLoadBalancer 获取CandidateL3NetworksForLoadBalancer详情
func (cli *ZSClient) GetCandidateL3NetworksForLoadBalancer(uuid string) (*view.GetCandidateL3NetworksForLoadBalancerView, error) {
	var resp view.GetCandidateL3NetworksForLoadBalancerView
	if err := cli.Get("v1/load-balancers/listeners/{listenerUuid}/networks/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}


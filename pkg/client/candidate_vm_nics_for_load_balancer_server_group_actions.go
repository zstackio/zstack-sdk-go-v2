// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateVmNicsForLoadBalancerServerGroup 获取CandidateVmNicsForLoadBalancerServerGroup详情
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancerServerGroup(uuid string) (*view.GetCandidateVmNicsForLoadBalancerServerGroupView, error) {
	var resp view.GetCandidateVmNicsForLoadBalancerServerGroupView
	if err := cli.Get("v1/load-balancers/servergroups/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}


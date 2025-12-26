// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateVmNicsForLoadBalancerServerGroup gets CandidateVmNicsForLoadBalancerServerGroup by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancerServerGroup(uuid string) (*view.GetCandidateVmNicsForLoadBalancerServerGroupView, error) {
	var resp view.GetCandidateVmNicsForLoadBalancerServerGroupView
	if err := cli.Get("v1/load-balancers/servergroups/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

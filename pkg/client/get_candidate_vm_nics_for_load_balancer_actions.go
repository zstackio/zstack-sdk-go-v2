// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateVmNicsForLoadBalancer gets CandidateVmNicsForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancer(uuid string) (*view.GetCandidateVmNicsForLoadBalancerView, error) {
	var resp view.GetCandidateVmNicsForLoadBalancerView
	if err := cli.Get("v1/load-balancers/listeners/{listenerUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

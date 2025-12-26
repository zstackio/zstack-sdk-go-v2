// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateL3NetworksForLoadBalancer gets CandidateL3NetworksForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateL3NetworksForLoadBalancer(uuid string) (*view.GetCandidateL3NetworksForLoadBalancerView, error) {
	var resp view.GetCandidateL3NetworksForLoadBalancerView
	if err := cli.Get("v1/load-balancers/listeners/{listenerUuid}/networks/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

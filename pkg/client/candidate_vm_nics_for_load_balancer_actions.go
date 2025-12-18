// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateVmNicsForLoadBalancer 获取CandidateVmNicsForLoadBalancer详情
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancer(uuid string) (*view.GetCandidateVmNicsForLoadBalancerView, error) {
	var resp view.GetCandidateVmNicsForLoadBalancerView
	if err := cli.Get("v1/load-balancers/listeners/{listenerUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}


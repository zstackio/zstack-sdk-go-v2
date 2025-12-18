// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateL3NetworksForChangeVmNicNetwork 获取CandidateL3NetworksForChangeVmNicNetwork详情
func (cli *ZSClient) GetCandidateL3NetworksForChangeVmNicNetwork(uuid string) (*view.GetCandidateL3NetworksForChangeVmNicNetworkView, error) {
	var resp view.GetCandidateL3NetworksForChangeVmNicNetworkView
	if err := cli.Get("v1/vm-instances/nics/{vmNicUuid}/l3-networks-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}


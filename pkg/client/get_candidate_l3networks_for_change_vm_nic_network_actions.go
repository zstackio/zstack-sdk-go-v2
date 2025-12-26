// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateL3NetworksForChangeVmNicNetwork gets CandidateL3NetworksForChangeVmNicNetwork by uuid
func (cli *ZSClient) GetCandidateL3NetworksForChangeVmNicNetwork(uuid string) (*view.GetCandidateL3NetworksForChangeVmNicNetworkView, error) {
	var resp view.GetCandidateL3NetworksForChangeVmNicNetworkView
	if err := cli.Get("v1/vm-instances/nics/{vmNicUuid}/l3-networks-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

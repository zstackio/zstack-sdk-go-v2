// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateNetworkInterfaces gets CandidateNetworkInterfaces by uuid
func (cli *ZSClient) GetCandidateNetworkInterfaces(uuid string) (*view.GetCandidateNetworkInterfacesView, error) {
	var resp view.GetCandidateNetworkInterfacesView
	if err := cli.Get("v1/cluster/hosts-network-interfaces", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

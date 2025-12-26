// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateInterfaceVlanIds gets CandidateInterfaceVlanIds by uuid
func (cli *ZSClient) GetCandidateInterfaceVlanIds(uuid string) (*view.GetCandidateInterfaceVlanIdsView, error) {
	var resp view.GetCandidateInterfaceVlanIdsView
	if err := cli.Get("v1/host/network-interface-vlan-ids", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

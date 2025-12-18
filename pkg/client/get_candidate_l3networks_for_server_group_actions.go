// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateL3NetworksForServerGroup gets CandidateL3NetworksForServerGroup by uuid
func (cli *ZSClient) GetCandidateL3NetworksForServerGroup(uuid string) (*view.GetCandidateL3NetworksForServerGroupView, error) {
	var resp view.GetCandidateL3NetworksForServerGroupView
	if err := cli.Get("v1/load-balancers/servergroups/candidate-l3network", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

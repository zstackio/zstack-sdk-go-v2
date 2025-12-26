// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateVmNicForSecurityGroup gets CandidateVmNicForSecurityGroup by uuid
func (cli *ZSClient) GetCandidateVmNicForSecurityGroup(uuid string) (*view.GetCandidateVmNicForSecurityGroupView, error) {
	var resp view.GetCandidateVmNicForSecurityGroupView
	if err := cli.Get("v1/security-groups/{securityGroupUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

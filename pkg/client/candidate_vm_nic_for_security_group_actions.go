// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateVmNicForSecurityGroup 获取CandidateVmNicForSecurityGroup详情
func (cli *ZSClient) GetCandidateVmNicForSecurityGroup(uuid string) (*view.GetCandidateVmNicForSecurityGroupView, error) {
	var resp view.GetCandidateVmNicForSecurityGroupView
	if err := cli.Get("v1/security-groups/{securityGroupUuid}/vm-instances/candidate-nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}


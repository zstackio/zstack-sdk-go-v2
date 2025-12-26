// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddVmNicToSecurityGroup adds VmNicToSecurityGroup
func (cli *ZSClient) AddVmNicToSecurityGroup(params param.AddVmNicToSecurityGroupParam) (*view.AddVmNicToSecurityGroupEventView, error) {
	resp := view.AddVmNicToSecurityGroupEventView{}
	if err := cli.Post("v1/security-groups/{securityGroupUuid}/vm-instances/nics", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

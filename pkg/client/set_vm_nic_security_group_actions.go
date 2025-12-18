// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmNicSecurityGroup operates on SetVmNicSecurityGroup
func (cli *ZSClient) SetVmNicSecurityGroup(uuid string, params param.SetVmNicSecurityGroupParam) (*view.SetVmNicSecurityGroupEventView, error) {
	resp := view.SetVmNicSecurityGroupEventView{}
	if err := cli.Put("v1/security-groups/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

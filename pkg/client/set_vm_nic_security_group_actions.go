// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmNicSecurityGroup 操作SetVmNicSecurityGroup
func (cli *ZSClient) SetVmNicSecurityGroup(uuid string, params param.SetVmNicSecurityGroupParam) (*view.SetVmNicSecurityGroupEventView, error) {
	resp := view.SetVmNicSecurityGroupEventView{}
	if err := cli.Put("v1/security-groups/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}


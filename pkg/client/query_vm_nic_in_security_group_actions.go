// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmNicInSecurityGroup queries VmNicInSecurityGroup list
func (cli *ZSClient) QueryVmNicInSecurityGroup(params *param.QueryParam) ([]view.VmNicSecurityGroupRefInventoryView, error) {
	var resp []view.VmNicSecurityGroupRefInventoryView
	return resp, cli.List("v1/security-groups/vm-instances/nics", params, &resp)
}

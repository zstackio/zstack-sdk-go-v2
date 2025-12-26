// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryVmNicSecurityPolicy queries VmNicSecurityPolicy list
func (cli *ZSClient) QueryVmNicSecurityPolicy(params *param.QueryParam) ([]view.VmNicSecurityPolicyInventoryView, error) {
	var resp []view.VmNicSecurityPolicyInventoryView
	return resp, cli.List("v1/security-groups/nics/security-policy", params, &resp)
}

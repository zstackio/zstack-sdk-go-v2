// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeVmNicSecurityPolicy changes VmNicSecurityPolicy
func (cli *ZSClient) ChangeVmNicSecurityPolicy(uuid string, params param.ChangeVmNicSecurityPolicyParam) (*view.VmNicSecurityPolicyInventoryView, error) {
	var resp view.ChangeVmNicSecurityPolicyEventView
	if err := cli.Put("v1/security-groups/nics/{vmNicUuid}/security-policy/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVmNicSecurityPolicy queries VmNicSecurityPolicy list
func (cli *ZSClient) QueryVmNicSecurityPolicy(params *param.QueryParam) ([]view.VmNicSecurityPolicyInventoryView, error) {
	var resp []view.VmNicSecurityPolicyInventoryView
	return resp, cli.List("v1/security-groups/nics/security-policy", params, &resp)
}

// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeVmNicSecurityPolicy changes VmNicSecurityPolicy
func (cli *ZSClient) ChangeVmNicSecurityPolicy(vmNicUuid string, params param.ChangeVmNicSecurityPolicyParam) (*view.VmNicSecurityPolicyInventoryView, error) {
	var resp view.ChangeVmNicSecurityPolicyEventView
	err := cli.PutWithSpec("v1/security-groups/nics", fmt.Sprintf(\"%s/security-policy/actions\", vmNicUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVmNicSecurityPolicy queries VmNicSecurityPolicy list
func (cli *ZSClient) QueryVmNicSecurityPolicy(params *param.QueryParam) ([]view.VmNicSecurityPolicyInventoryView, error) {
	var resp []view.VmNicSecurityPolicyInventoryView
	return resp, cli.List("v1/security-groups/nics/security-policy", params, &resp)
}

func (cli *ZSClient) GetVmNicSecurityPolicy(uuid string) (*view.VmNicSecurityPolicyInventoryView, error) {
	var resp view.VmNicSecurityPolicyInventoryView
	if err := cli.Get("v1/security-groups/nics/security-policy", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

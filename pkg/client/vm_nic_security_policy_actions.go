// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeVmNicSecurityPolicy changes VmNicSecurityPolicy
func (cli *ZSClient) ChangeVmNicSecurityPolicy(ctx context.Context, vmNicUuid string, params param.ChangeVmNicSecurityPolicyParam) (*view.VmNicSecurityPolicyInventoryView, error) {
	resp := view.VmNicSecurityPolicyInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-groups/nics", vmNicUuid, "", map[string]interface{}{
		"changeVmNicSecurityPolicy": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVmNicSecurityPolicy queries VmNicSecurityPolicy list
func (cli *ZSClient) QueryVmNicSecurityPolicy(ctx context.Context, params *param.QueryParam) ([]view.VmNicSecurityPolicyInventoryView, error) {
	var resp []view.VmNicSecurityPolicyInventoryView
	return resp, cli.List(ctx, "v1/security-groups/nics/security-policy", params, &resp)
}

func (cli *ZSClient) GetVmNicSecurityPolicy(ctx context.Context, uuid string) (*view.VmNicSecurityPolicyInventoryView, error) {
	var resp view.VmNicSecurityPolicyInventoryView
	if err := cli.Get(ctx, "v1/security-groups/nics/security-policy", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmNicSecurityPolicy Pagination
func (cli *ZSClient) PageVmNicSecurityPolicy(ctx context.Context, params *param.QueryParam) ([]view.VmNicSecurityPolicyInventoryView, int, error) {
	var vmNicSecurityPolicies []view.VmNicSecurityPolicyInventoryView
	total, err := cli.Page(ctx, "v1/security-groups/nics/security-policy", params, &vmNicSecurityPolicies)
	return vmNicSecurityPolicies, total, err
}

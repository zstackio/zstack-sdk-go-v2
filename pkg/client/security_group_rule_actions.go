// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeSecurityGroupRule changes SecurityGroupRule
func (cli *ZSClient) ChangeSecurityGroupRule(ctx context.Context, uuid string, params param.ChangeSecurityGroupRuleParam) (*view.SecurityGroupRuleInventoryView, error) {
	resp := view.SecurityGroupRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-groups/rules", uuid, "", map[string]interface{}{
		"changeSecurityGroupRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ValidateSecurityGroupRule operates on SecurityGroupRule
func (cli *ZSClient) ValidateSecurityGroupRule(ctx context.Context, uuid string) (*view.SecurityGroupRuleInventoryView, error) {
	var resp view.SecurityGroupRuleInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/security-groups", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddSecurityGroupRule adds SecurityGroupRule
func (cli *ZSClient) AddSecurityGroupRule(ctx context.Context, params param.AddSecurityGroupRuleParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.Post(ctx, "v1/security-groups/{securityGroupUuid}/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySecurityGroupRule queries SecurityGroupRule list
func (cli *ZSClient) QuerySecurityGroupRule(ctx context.Context, params *param.QueryParam) ([]view.SecurityGroupRuleInventoryView, error) {
	var resp []view.SecurityGroupRuleInventoryView
	return resp, cli.List(ctx, "v1/security-groups/rules", params, &resp)
}

func (cli *ZSClient) GetSecurityGroupRule(ctx context.Context, uuid string) (*view.SecurityGroupRuleInventoryView, error) {
	var resp view.SecurityGroupRuleInventoryView
	if err := cli.Get(ctx, "v1/security-groups/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSecurityGroupRule Pagination
func (cli *ZSClient) PageSecurityGroupRule(ctx context.Context, params *param.QueryParam) ([]view.SecurityGroupRuleInventoryView, int, error) {
	var securityGroupRules []view.SecurityGroupRuleInventoryView
	total, err := cli.Page(ctx, "v1/security-groups/rules", params, &securityGroupRules)
	return securityGroupRules, total, err
}
// DeleteSecurityGroupRule deletes SecurityGroupRule
func (cli *ZSClient) DeleteSecurityGroupRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/security-groups/rules", uuid, string(deleteMode))
}

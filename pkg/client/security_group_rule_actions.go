// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeSecurityGroupRule changes SecurityGroupRule
func (cli *ZSClient) ChangeSecurityGroupRule(uuid string, params param.ChangeSecurityGroupRuleParam) (*view.SecurityGroupRuleInventoryView, error) {
	var resp view.ChangeSecurityGroupRuleEventView
	if err := cli.Put("v1/security-groups/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ValidateSecurityGroupRule operates on SecurityGroupRule
func (cli *ZSClient) ValidateSecurityGroupRule(params param.ValidateSecurityGroupRuleParam) (*view.SecurityGroupRuleInventoryView, error) {
	var resp view.SecurityGroupRuleInventoryView
	if err := cli.Get("v1/security-groups/{securityGroupUuid}/rules/validation", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddSecurityGroupRule adds SecurityGroupRule
func (cli *ZSClient) AddSecurityGroupRule(params param.AddSecurityGroupRuleParam) (*view.SecurityGroupInventoryView, error) {
	var resp view.AddSecurityGroupRuleEventView
	if err := cli.Post("v1/security-groups/{securityGroupUuid}/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySecurityGroupRule queries SecurityGroupRule list
func (cli *ZSClient) QuerySecurityGroupRule(params *param.QueryParam) ([]view.SecurityGroupRuleInventoryView, error) {
	var resp []view.SecurityGroupRuleInventoryView
	return resp, cli.List("v1/security-groups/rules", params, &resp)
}

func (cli *ZSClient) GetSecurityGroupRule(uuid string) (*view.SecurityGroupRuleInventoryView, error) {
	var resp view.SecurityGroupRuleInventoryView
	if err := cli.Get("v1/security-groups/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSecurityGroupRule deletes SecurityGroupRule
func (cli *ZSClient) DeleteSecurityGroupRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups/rules", uuid, string(deleteMode))
}

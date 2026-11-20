// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ChangeSecurityGroupRule changes SecurityGroupRule
func (cli *ZSClient) ChangeSecurityGroupRule(uuid string, params param.ChangeSecurityGroupRuleParam) (*view.SecurityGroupRuleInventoryView, error) {
	resp := view.SecurityGroupRuleInventoryView{}
	if err := cli.PutWithRespKey("v1/security-groups/rules", uuid, "", map[string]interface{}{
		"changeSecurityGroupRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ValidateSecurityGroupRule operates on SecurityGroupRule
func (cli *ZSClient) ValidateSecurityGroupRule(uuid string) (*view.SecurityGroupRuleInventoryView, error) {
	var resp view.SecurityGroupRuleInventoryView
	if err := cli.GetWithRespKey("v1/security-groups", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddSecurityGroupRule adds SecurityGroupRule
func (cli *ZSClient) AddSecurityGroupRule(securityGroupUuid string, params param.AddSecurityGroupRuleParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/security-groups/%s/rules", securityGroupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageSecurityGroupRule Pagination
func (cli *ZSClient) PageSecurityGroupRule(params *param.QueryParam) ([]view.SecurityGroupRuleInventoryView, int, error) {
	var securityGroupRules []view.SecurityGroupRuleInventoryView
	total, err := cli.Page("v1/security-groups/rules", params, &securityGroupRules)
	return securityGroupRules, total, err
}
// DeleteSecurityGroupRule deletes one security group rule.
func (cli *ZSClient) DeleteSecurityGroupRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteSecurityGroupRuleWithParams(deleteMode, param.DeleteSecurityGroupRuleParamDetail{RuleUuids: []string{uuid}})
}

// DeleteSecurityGroupRuleWithParams sends the required route and query parameters for DeleteSecurityGroupRule.
func (cli *ZSClient) DeleteSecurityGroupRuleWithParams(deleteMode param.DeleteMode, params param.DeleteSecurityGroupRuleParamDetail) error {
	query := url.Values{
		"ruleUuids": params.RuleUuids,
	}
	return cli.deleteActionWithQuery("v1/security-groups/rules", "", "", deleteMode, query, "ruleUuids")
}

// ChangeSecurityGroupRuleState changes SecurityGroupRuleState
func (cli *ZSClient) ChangeSecurityGroupRuleState(securityGroupUuid string, params param.ChangeSecurityGroupRuleStateParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/security-groups", securityGroupUuid, "", map[string]interface{}{
		"changeSecurityGroupRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecurityGroupRulePriority updates SecurityGroupRulePriority
func (cli *ZSClient) UpdateSecurityGroupRulePriority(securityGroupUuid string, params param.UpdateSecurityGroupRulePriorityParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/security-groups", securityGroupUuid, "", map[string]interface{}{
		"updateSecurityGroupRulePriority": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

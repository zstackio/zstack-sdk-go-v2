// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// ChangeFirewallRuleState changes FirewallRuleState
func (cli *ZSClient) ChangeFirewallRuleState(uuid string, params param.ChangeFirewallRuleStateParam) (*view.VpcFirewallRuleInventoryView, error) {
	resp := view.VpcFirewallRuleInventoryView{}
	if err := cli.PutWithRespKey("v1/vpcfirewalls/rules", uuid, "", map[string]interface{}{
		"changeFirewallRuleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRuleFromConfigFile creates FirewallRuleFromConfigFile
func (cli *ZSClient) CreateFirewallRuleFromConfigFile(params param.CreateFirewallRuleFromConfigFileParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.Post("v1/vpcfirewalls/rules/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckFirewallRuleConfigFile operates on FirewallRuleConfigFile
func (cli *ZSClient) CheckFirewallRuleConfigFile(params param.CheckFirewallRuleConfigFileParam) (*view.CheckFirewallRuleConfigFileView, error) {
	resp := view.CheckFirewallRuleConfigFileView{}
	if err := cli.Post("v1/vpcfirewalls/rules/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFirewallRule creates FirewallRule
func (cli *ZSClient) CreateFirewallRule(params param.CreateFirewallRuleParam) (*view.VpcFirewallRuleInventoryView, error) {
	resp := view.VpcFirewallRuleInventoryView{}
	if err := cli.Post("v1/vpcfirewalls/rules", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewallRule deletes FirewallRule
func (cli *ZSClient) DeleteFirewallRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/rules", uuid, string(deleteMode))
}

// QueryFirewallRule queries FirewallRule list
func (cli *ZSClient) QueryFirewallRule(params *param.QueryParam) ([]view.VpcFirewallRuleInventoryView, error) {
	var resp []view.VpcFirewallRuleInventoryView
	return resp, cli.List("v1/vpcfirewalls/rules", params, &resp)
}

func (cli *ZSClient) GetFirewallRule(uuid string) (*view.VpcFirewallRuleInventoryView, error) {
	var resp view.VpcFirewallRuleInventoryView
	if err := cli.Get("v1/vpcfirewalls/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRule Pagination
func (cli *ZSClient) PageFirewallRule(params *param.QueryParam) ([]view.VpcFirewallRuleInventoryView, int, error) {
	var firewallRules []view.VpcFirewallRuleInventoryView
	total, err := cli.Page("v1/vpcfirewalls/rules", params, &firewallRules)
	return firewallRules, total, err
}

// UpdateFirewallRule updates FirewallRule
func (cli *ZSClient) UpdateFirewallRule(uuid string, params param.UpdateFirewallRuleParam) (*view.VpcFirewallRuleInventoryView, error) {
	resp := view.VpcFirewallRuleInventoryView{}
	if err := cli.PutWithRespKey("v1/vpcfirewalls/rules", uuid, "", map[string]interface{}{
		"updateFirewallRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

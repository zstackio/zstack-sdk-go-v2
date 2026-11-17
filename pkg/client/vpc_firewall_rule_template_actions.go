// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// QueryFirewallRuleTemplate queries FirewallRuleTemplate list
func (cli *ZSClient) QueryFirewallRuleTemplate(params *param.QueryParam) ([]view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp []view.VpcFirewallRuleTemplateInventoryView
	return resp, cli.List("v1/vpcfirewalls/rules/templates", params, &resp)
}

func (cli *ZSClient) GetFirewallRuleTemplate(uuid string) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	var resp view.VpcFirewallRuleTemplateInventoryView
	if err := cli.Get("v1/vpcfirewalls/rules/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRuleTemplate Pagination
func (cli *ZSClient) PageFirewallRuleTemplate(params *param.QueryParam) ([]view.VpcFirewallRuleTemplateInventoryView, int, error) {
	var firewallRuleTemplates []view.VpcFirewallRuleTemplateInventoryView
	total, err := cli.Page("v1/vpcfirewalls/rules/templates", params, &firewallRuleTemplates)
	return firewallRuleTemplates, total, err
}

// CreateFirewallRuleTemplate creates FirewallRuleTemplate
func (cli *ZSClient) CreateFirewallRuleTemplate(params param.CreateFirewallRuleTemplateParam) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	resp := view.VpcFirewallRuleTemplateInventoryView{}
	if err := cli.Post("v1/vpcfirewalls/rules/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFirewallRuleTemplate deletes FirewallRuleTemplate
func (cli *ZSClient) DeleteFirewallRuleTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/rules/templates", uuid, string(deleteMode))
}

// UpdateFirewallRuleTemplate updates FirewallRuleTemplate
func (cli *ZSClient) UpdateFirewallRuleTemplate(uuid string, params param.UpdateFirewallRuleTemplateParam) (*view.VpcFirewallRuleTemplateInventoryView, error) {
	resp := view.VpcFirewallRuleTemplateInventoryView{}
	if err := cli.PutWithRespKey("v1/vpcfirewalls/rules/template", uuid, "", map[string]interface{}{
		"updateFirewallRuleTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

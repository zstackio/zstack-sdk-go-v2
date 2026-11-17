// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DeleteFirewallIpSetTemplate deletes FirewallIpSetTemplate
func (cli *ZSClient) DeleteFirewallIpSetTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/ipset/templates", uuid, string(deleteMode))
}

// CreateFirewallIpSetTemplate creates FirewallIpSetTemplate
func (cli *ZSClient) CreateFirewallIpSetTemplate(params param.CreateFirewallIpSetTemplateParam) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	resp := view.VpcFirewallIpSetTemplateInventoryView{}
	if err := cli.Post("v1/vpcfirewalls/ipset/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallIpSetTemplate queries FirewallIpSetTemplate list
func (cli *ZSClient) QueryFirewallIpSetTemplate(params *param.QueryParam) ([]view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp []view.VpcFirewallIpSetTemplateInventoryView
	return resp, cli.List("v1/vpcfirewalls/ipset/templates", params, &resp)
}

func (cli *ZSClient) GetFirewallIpSetTemplate(uuid string) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	var resp view.VpcFirewallIpSetTemplateInventoryView
	if err := cli.Get("v1/vpcfirewalls/ipset/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallIpSetTemplate Pagination
func (cli *ZSClient) PageFirewallIpSetTemplate(params *param.QueryParam) ([]view.VpcFirewallIpSetTemplateInventoryView, int, error) {
	var firewallIpSetTemplates []view.VpcFirewallIpSetTemplateInventoryView
	total, err := cli.Page("v1/vpcfirewalls/ipset/templates", params, &firewallIpSetTemplates)
	return firewallIpSetTemplates, total, err
}

// UpdateFirewallIpSetTemplate updates FirewallIpSetTemplate
func (cli *ZSClient) UpdateFirewallIpSetTemplate(uuid string, params param.UpdateFirewallIpSetTemplateParam) (*view.VpcFirewallIpSetTemplateInventoryView, error) {
	resp := view.VpcFirewallIpSetTemplateInventoryView{}
	if err := cli.PutWithRespKey("v1/vpcfirewalls/ipset/templates", uuid, "", map[string]interface{}{
		"updateFirewallIpSetTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

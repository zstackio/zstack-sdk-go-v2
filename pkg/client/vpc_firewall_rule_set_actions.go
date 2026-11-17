// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DetachFirewallRuleSetFromL3 operates on FirewallRuleSetFromL3
func (cli *ZSClient) DetachFirewallRuleSetFromL3(l3Uuid string, ruleSetUuid string, params param.DetachFirewallRuleSetFromL3Param) (*view.DetachFirewallRuleSetFromL3EventView, error) {
	resp := view.DetachFirewallRuleSetFromL3EventView{}
	if err := cli.Post(fmt.Sprintf("v1/vpcfirewalls/l3networks/%s/ruleSets/%s", l3Uuid, ruleSetUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachFirewallRuleSetToL3 operates on FirewallRuleSetToL3
func (cli *ZSClient) AttachFirewallRuleSetToL3(ruleSetUuid string, l3Uuid string, params param.AttachFirewallRuleSetToL3Param) (*view.VpcFirewallRuleSetL3RefInventoryView, error) {
	resp := view.VpcFirewallRuleSetL3RefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vpcfirewalls/ruleSets/%s/l3networks/%s", ruleSetUuid, l3Uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFirewallRuleSet queries FirewallRuleSet list
func (cli *ZSClient) QueryFirewallRuleSet(params *param.QueryParam) ([]view.VpcFirewallRuleSetInventoryView, error) {
	var resp []view.VpcFirewallRuleSetInventoryView
	return resp, cli.List("v1/vpcfirewalls/ruleSets", params, &resp)
}

func (cli *ZSClient) GetFirewallRuleSet(uuid string) (*view.VpcFirewallRuleSetInventoryView, error) {
	var resp view.VpcFirewallRuleSetInventoryView
	if err := cli.Get("v1/vpcfirewalls/ruleSets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFirewallRuleSet Pagination
func (cli *ZSClient) PageFirewallRuleSet(params *param.QueryParam) ([]view.VpcFirewallRuleSetInventoryView, int, error) {
	var firewallRuleSets []view.VpcFirewallRuleSetInventoryView
	total, err := cli.Page("v1/vpcfirewalls/ruleSets", params, &firewallRuleSets)
	return firewallRuleSets, total, err
}

// DeleteFirewallRuleSet deletes FirewallRuleSet
func (cli *ZSClient) DeleteFirewallRuleSet(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vpcfirewalls/ruleSets", uuid, string(deleteMode))
}

// CreateFirewallRuleSet creates FirewallRuleSet
func (cli *ZSClient) CreateFirewallRuleSet(params param.CreateFirewallRuleSetParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.Post("v1/vpcfirewalls/ruleSets", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ApplyRuleSetChanges operates on RuleSetChanges
func (cli *ZSClient) ApplyRuleSetChanges(uuid string, params param.ApplyRuleSetChangesParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.PutWithRespKey("v1/vpcfirewalls/ruleSets/apply", uuid, "", map[string]interface{}{
		"applyRuleSetChanges": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFirewallRuleSet updates FirewallRuleSet
func (cli *ZSClient) UpdateFirewallRuleSet(uuid string, params param.UpdateFirewallRuleSetParam) (*view.VpcFirewallRuleSetInventoryView, error) {
	resp := view.VpcFirewallRuleSetInventoryView{}
	if err := cli.PutWithRespKey("v1/vpcfirewalls/ruleSets", uuid, "", map[string]interface{}{
		"updateFirewallRuleSet": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

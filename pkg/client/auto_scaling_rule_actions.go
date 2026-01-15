// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAutoScalingRule updates AutoScalingRule
func (cli *ZSClient) UpdateAutoScalingRule(uuid string, params param.UpdateAutoScalingRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.Put("v1/autoscaling/rules", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAutoScalingRule deletes AutoScalingRule
func (cli *ZSClient) DeleteAutoScalingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/rules", uuid, string(deleteMode))
}
// QueryAutoScalingRule queries AutoScalingRule list
func (cli *ZSClient) QueryAutoScalingRule(params *param.QueryParam) ([]view.AutoScalingRuleInventoryView, error) {
	var resp []view.AutoScalingRuleInventoryView
	return resp, cli.List("v1/autoscaling/groups/rules", params, &resp)
}

func (cli *ZSClient) GetAutoScalingRule(uuid string) (*view.AutoScalingRuleInventoryView, error) {
	var resp view.AutoScalingRuleInventoryView
	if err := cli.Get("v1/autoscaling/groups/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAutoScalingRule Pagination
func (cli *ZSClient) PageAutoScalingRule(params *param.QueryParam) ([]view.AutoScalingRuleInventoryView, int, error) {
	var autoScalingRules []view.AutoScalingRuleInventoryView
	total, err := cli.Page("v1/autoscaling/groups/rules", params, &autoScalingRules)
	return autoScalingRules, total, err
}

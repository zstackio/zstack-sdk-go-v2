// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAutoScalingRule updates AutoScalingRule
func (cli *ZSClient) UpdateAutoScalingRule(uuid string, params param.UpdateAutoScalingRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	var resp view.UpdateAutoScalingRuleEventView
	if err := cli.Put("v1/autoscaling/rules/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAutoScalingRule deletes AutoScalingRule
func (cli *ZSClient) DeleteAutoScalingRule(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/rules/{uuid}", uuid, string(deleteMode))
}
// QueryAutoScalingRule queries AutoScalingRule list
func (cli *ZSClient) QueryAutoScalingRule(params *param.QueryParam) ([]view.AutoScalingRuleInventoryView, error) {
	var resp []view.AutoScalingRuleInventoryView
	return resp, cli.List("v1/autoscaling/groups/rules", params, &resp)
}

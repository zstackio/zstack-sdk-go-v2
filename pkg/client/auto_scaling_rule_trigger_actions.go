// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAutoScalingRuleTrigger queries AutoScalingRuleTrigger list
func (cli *ZSClient) QueryAutoScalingRuleTrigger(params *param.QueryParam) ([]view.AutoScalingRuleTriggerInventoryView, error) {
	var resp []view.AutoScalingRuleTriggerInventoryView
	return resp, cli.List("v1/autoscaling/groups/rules/trigger", params, &resp)
}

// PageAutoScalingRuleTrigger Pagination
func (cli *ZSClient) PageAutoScalingRuleTrigger(params *param.QueryParam) ([]view.AutoScalingRuleTriggerInventoryView, int, error) {
	var autoScalingRuleTriggers []view.AutoScalingRuleTriggerInventoryView
	total, err := cli.Page("v1/autoscaling/groups/rules/trigger", params, &autoScalingRuleTriggers)
	return autoScalingRuleTriggers, total, err
}
// DeleteAutoScalingRuleTrigger deletes AutoScalingRuleTrigger
func (cli *ZSClient) DeleteAutoScalingRuleTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/rules/triggers", uuid, string(deleteMode))
}

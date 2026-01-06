// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAutoScalingRuleTrigger queries AutoScalingRuleTrigger list
func (cli *ZSClient) QueryAutoScalingRuleTrigger(params *param.QueryParam) ([]view.AutoScalingRuleTriggerInventoryView, error) {
	var resp []view.AutoScalingRuleTriggerInventoryView
	return resp, cli.List("v1/autoscaling/groups/rules/trigger", params, &resp)
}
// DeleteAutoScalingRuleTrigger deletes AutoScalingRuleTrigger
func (cli *ZSClient) DeleteAutoScalingRuleTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/groups/rules/triggers/{uuid}", uuid, string(deleteMode))
}

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

func (cli *ZSClient) GetAutoScalingRuleTrigger(uuid string) (*view.AutoScalingRuleTriggerInventoryView, error) {
	var resp view.AutoScalingRuleTriggerInventoryView
	if err := cli.Get("v1/autoscaling/groups/rules/trigger", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAutoScalingRuleTrigger deletes AutoScalingRuleTrigger
func (cli *ZSClient) DeleteAutoScalingRuleTrigger(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/autoscaling/groups/rules/triggers", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}

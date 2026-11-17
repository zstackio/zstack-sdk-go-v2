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
	if err := cli.PutWithRespKey("v1/autoscaling/rules", uuid, "", map[string]interface{}{
		"updateAutoScalingRule": params.Params,
	}, &resp); err != nil {
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

// CreateAutoScalingGroupAddingNewInstanceRule creates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupAddingNewInstanceRule(params param.CreateAutoScalingGroupAddingNewInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.Post("v1/autoscaling/rules/adding-new-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAutoScalingGroupRemovalInstanceRule creates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) CreateAutoScalingGroupRemovalInstanceRule(params param.CreateAutoScalingGroupRemovalInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.Post("v1/autoscaling/rules/removal-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAutoScalingGroupRemovalInstanceRule updates AutoScalingGroupRemovalInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupRemovalInstanceRule(uuid string, params param.UpdateAutoScalingGroupRemovalInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.PutWithRespKey("v1/autoscaling/rules/removal-instance", uuid, "", map[string]interface{}{
		"updateAutoScalingGroupRemovalInstanceRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExecuteAutoScalingRule operates on ExecuteAutoScalingRule
func (cli *ZSClient) ExecuteAutoScalingRule(uuid string, params param.ExecuteAutoScalingRuleParam) (*view.ExecuteAutoScalingRuleEventView, error) {
	resp := view.ExecuteAutoScalingRuleEventView{}
	if err := cli.PutWithRespKey("v1/autoscaling/rules", uuid, "", map[string]interface{}{
		"executeAutoScalingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAutoScalingGroupAddingNewInstanceRule updates AutoScalingGroupAddingNewInstanceRule
func (cli *ZSClient) UpdateAutoScalingGroupAddingNewInstanceRule(uuid string, params param.UpdateAutoScalingGroupAddingNewInstanceRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.PutWithRespKey("v1/autoscaling/rules/adding-new-instance", uuid, "", map[string]interface{}{
		"updateAutoScalingGroupAddingNewInstanceRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

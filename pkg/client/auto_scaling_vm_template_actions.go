// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAutoScalingVmTemplate updates AutoScalingVmTemplate
func (cli *ZSClient) UpdateAutoScalingVmTemplate(uuid string, params param.UpdateAutoScalingVmTemplateParam) (*view.AutoScalingTemplateInventoryView, error) {
	var resp view.UpdateAutoScalingTemplateEventView
	if err := cli.Put("v1/autoscaling/vmtemplate/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryAutoScalingVmTemplate queries AutoScalingVmTemplate list
func (cli *ZSClient) QueryAutoScalingVmTemplate(params *param.QueryParam) ([]view.AutoScalingVmTemplateInventoryView, error) {
	var resp []view.AutoScalingVmTemplateInventoryView
	return resp, cli.List("v1/autoscaling/vmtemplate", params, &resp)
}

func (cli *ZSClient) GetAutoScalingVmTemplate(uuid string) (*view.AutoScalingVmTemplateInventoryView, error) {
	var resp view.AutoScalingVmTemplateInventoryView
	if err := cli.Get("v1/autoscaling/vmtemplate", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateAutoScalingVmTemplate creates AutoScalingVmTemplate
func (cli *ZSClient) CreateAutoScalingVmTemplate(params param.CreateAutoScalingVmTemplateParam) (*view.AutoScalingTemplateInventoryView, error) {
	var resp view.CreateAutoScalingTemplateEventView
	if err := cli.Post("v1/autoscaling/vmtemplate", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

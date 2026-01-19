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
	resp := view.AutoScalingTemplateInventoryView{}
	if err := cli.Put("v1/autoscaling/vmtemplate", uuid, map[string]interface{}{
		"updateAutoScalingVmTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageAutoScalingVmTemplate Pagination
func (cli *ZSClient) PageAutoScalingVmTemplate(params *param.QueryParam) ([]view.AutoScalingVmTemplateInventoryView, int, error) {
	var autoScalingVmTemplates []view.AutoScalingVmTemplateInventoryView
	total, err := cli.Page("v1/autoscaling/vmtemplate", params, &autoScalingVmTemplates)
	return autoScalingVmTemplates, total, err
}
// CreateAutoScalingVmTemplate creates AutoScalingVmTemplate
func (cli *ZSClient) CreateAutoScalingVmTemplate(params param.CreateAutoScalingVmTemplateParam) (*view.AutoScalingTemplateInventoryView, error) {
	resp := view.AutoScalingTemplateInventoryView{}
	if err := cli.Post("v1/autoscaling/vmtemplate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

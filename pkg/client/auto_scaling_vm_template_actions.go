// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAutoScalingVmTemplate updates AutoScalingVmTemplate
func (cli *ZSClient) UpdateAutoScalingVmTemplate(ctx context.Context, uuid string, params param.UpdateAutoScalingVmTemplateParam) (*view.AutoScalingTemplateInventoryView, error) {
	resp := view.AutoScalingTemplateInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/autoscaling/vmtemplate", uuid, "", map[string]interface{}{
		"updateAutoScalingVmTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAutoScalingVmTemplate queries AutoScalingVmTemplate list
func (cli *ZSClient) QueryAutoScalingVmTemplate(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingVmTemplateInventoryView, error) {
	var resp []view.AutoScalingVmTemplateInventoryView
	return resp, cli.List(ctx, "v1/autoscaling/vmtemplate", params, &resp)
}

func (cli *ZSClient) GetAutoScalingVmTemplate(ctx context.Context, uuid string) (*view.AutoScalingVmTemplateInventoryView, error) {
	var resp view.AutoScalingVmTemplateInventoryView
	if err := cli.Get(ctx, "v1/autoscaling/vmtemplate", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAutoScalingVmTemplate Pagination
func (cli *ZSClient) PageAutoScalingVmTemplate(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingVmTemplateInventoryView, int, error) {
	var autoScalingVmTemplates []view.AutoScalingVmTemplateInventoryView
	total, err := cli.Page(ctx, "v1/autoscaling/vmtemplate", params, &autoScalingVmTemplates)
	return autoScalingVmTemplates, total, err
}
// CreateAutoScalingVmTemplate creates AutoScalingVmTemplate
func (cli *ZSClient) CreateAutoScalingVmTemplate(ctx context.Context, params param.CreateAutoScalingVmTemplateParam) (*view.AutoScalingTemplateInventoryView, error) {
	resp := view.AutoScalingTemplateInventoryView{}
	if err := cli.Post(ctx, "v1/autoscaling/vmtemplate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

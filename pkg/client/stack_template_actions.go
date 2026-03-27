// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryStackTemplate queries StackTemplate list
func (cli *ZSClient) QueryStackTemplate(ctx context.Context, params *param.QueryParam) ([]view.StackTemplateInventoryView, error) {
	var resp []view.StackTemplateInventoryView
	return resp, cli.List(ctx, "v1/cloudformation/template", params, &resp)
}

func (cli *ZSClient) GetStackTemplate(ctx context.Context, uuid string) (*view.StackTemplateInventoryView, error) {
	var resp view.StackTemplateInventoryView
	if err := cli.Get(ctx, "v1/cloudformation/template", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageStackTemplate Pagination
func (cli *ZSClient) PageStackTemplate(ctx context.Context, params *param.QueryParam) ([]view.StackTemplateInventoryView, int, error) {
	var stackTemplates []view.StackTemplateInventoryView
	total, err := cli.Page(ctx, "v1/cloudformation/template", params, &stackTemplates)
	return stackTemplates, total, err
}
// DeleteStackTemplate deletes StackTemplate
func (cli *ZSClient) DeleteStackTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/cloudformation/template", uuid, string(deleteMode))
}
// UpdateStackTemplate updates StackTemplate
func (cli *ZSClient) UpdateStackTemplate(ctx context.Context, uuid string, params param.UpdateStackTemplateParam) (*view.StackTemplateInventoryView, error) {
	resp := view.StackTemplateInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/cloudformation/template", uuid, "", map[string]interface{}{
		"updateStackTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddStackTemplate adds StackTemplate
func (cli *ZSClient) AddStackTemplate(ctx context.Context, params param.AddStackTemplateParam) (*view.StackTemplateInventoryView, error) {
	resp := view.StackTemplateInventoryView{}
	if err := cli.Post(ctx, "v1/cloudformation/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

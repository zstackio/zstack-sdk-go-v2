// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryStackTemplate queries StackTemplate list
func (cli *ZSClient) QueryStackTemplate(params *param.QueryParam) ([]view.StackTemplateInventoryView, error) {
	var resp []view.StackTemplateInventoryView
	return resp, cli.List("v1/cloudformation/template", params, &resp)
}

func (cli *ZSClient) GetStackTemplate(uuid string) (*view.StackTemplateInventoryView, error) {
	var resp view.StackTemplateInventoryView
	if err := cli.Get("v1/cloudformation/template", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageStackTemplate Pagination
func (cli *ZSClient) PageStackTemplate(params *param.QueryParam) ([]view.StackTemplateInventoryView, int, error) {
	var stackTemplates []view.StackTemplateInventoryView
	total, err := cli.Page("v1/cloudformation/template", params, &stackTemplates)
	return stackTemplates, total, err
}
// DeleteStackTemplate deletes StackTemplate
func (cli *ZSClient) DeleteStackTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/template", uuid, string(deleteMode))
}
// UpdateStackTemplate updates StackTemplate
func (cli *ZSClient) UpdateStackTemplate(uuid string, params param.UpdateStackTemplateParam) (*view.StackTemplateInventoryView, error) {
	resp := view.StackTemplateInventoryView{}
	if err := cli.Put("v1/cloudformation/template", uuid, map[string]interface{}{
		"updateStackTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddStackTemplate adds StackTemplate
func (cli *ZSClient) AddStackTemplate(params param.AddStackTemplateParam) (*view.StackTemplateInventoryView, error) {
	resp := view.StackTemplateInventoryView{}
	if err := cli.Post("v1/cloudformation/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

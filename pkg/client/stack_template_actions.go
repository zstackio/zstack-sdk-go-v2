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
// DeleteStackTemplate deletes StackTemplate
func (cli *ZSClient) DeleteStackTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/cloudformation/template", uuid, string(deleteMode))
}
// UpdateStackTemplate updates StackTemplate
func (cli *ZSClient) UpdateStackTemplate(uuid string, params param.UpdateStackTemplateParam) (*view.StackTemplateInventoryView, error) {
	var resp view.UpdateStackTemplateEventView
	if err := cli.Put("v1/cloudformation/template/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// AddStackTemplate adds StackTemplate
func (cli *ZSClient) AddStackTemplate(params param.AddStackTemplateParam) (*view.StackTemplateInventoryView, error) {
	var resp view.AddStackTemplateEventView
	if err := cli.Post("v1/cloudformation/template", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

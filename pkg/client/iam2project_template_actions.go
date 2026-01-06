// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateIAM2ProjectTemplate creates IAM2ProjectTemplate
func (cli *ZSClient) CreateIAM2ProjectTemplate(params param.CreateIAM2ProjectTemplateParam) (*view.IAM2ProjectTemplateInventoryView, error) {
	var resp view.CreateIAM2ProjectTemplateEventView
	if err := cli.Post("v1/iam2/projects/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryIAM2ProjectTemplate queries IAM2ProjectTemplate list
func (cli *ZSClient) QueryIAM2ProjectTemplate(params *param.QueryParam) ([]view.IAM2ProjectTemplateInventoryView, error) {
	var resp []view.IAM2ProjectTemplateInventoryView
	return resp, cli.List("v1/iam2/projects/templates", params, &resp)
}
// UpdateIAM2ProjectTemplate updates IAM2ProjectTemplate
func (cli *ZSClient) UpdateIAM2ProjectTemplate(uuid string, params param.UpdateIAM2ProjectTemplateParam) (*view.IAM2ProjectTemplateInventoryView, error) {
	var resp view.UpdateIAM2ProjectTemplateEventView
	if err := cli.Put("v1/iam2/projects/templates/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteIAM2ProjectTemplate deletes IAM2ProjectTemplate
func (cli *ZSClient) DeleteIAM2ProjectTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/templates/{uuid}", uuid, string(deleteMode))
}

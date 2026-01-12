// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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

func (cli *ZSClient) GetIAM2ProjectTemplate(uuid string) (*view.IAM2ProjectTemplateInventoryView, error) {
	var resp view.IAM2ProjectTemplateInventoryView
	if err := cli.Get("v1/iam2/projects/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIAM2ProjectTemplate updates IAM2ProjectTemplate
func (cli *ZSClient) UpdateIAM2ProjectTemplate(uuid string, params param.UpdateIAM2ProjectTemplateParam) (*view.IAM2ProjectTemplateInventoryView, error) {
	var resp view.UpdateIAM2ProjectTemplateEventView
	err := cli.PutWithSpec("v1/iam2/projects/templates", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteIAM2ProjectTemplate deletes IAM2ProjectTemplate
func (cli *ZSClient) DeleteIAM2ProjectTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/iam2/projects/templates", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}

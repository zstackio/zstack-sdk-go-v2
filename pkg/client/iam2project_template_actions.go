// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateIAM2ProjectTemplate creates IAM2ProjectTemplate
func (cli *ZSClient) CreateIAM2ProjectTemplate(ctx context.Context, params param.CreateIAM2ProjectTemplateParam) (*view.IAM2ProjectTemplateInventoryView, error) {
	resp := view.IAM2ProjectTemplateInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/projects/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIAM2ProjectTemplate queries IAM2ProjectTemplate list
func (cli *ZSClient) QueryIAM2ProjectTemplate(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectTemplateInventoryView, error) {
	var resp []view.IAM2ProjectTemplateInventoryView
	return resp, cli.List(ctx, "v1/iam2/projects/templates", params, &resp)
}

func (cli *ZSClient) GetIAM2ProjectTemplate(ctx context.Context, uuid string) (*view.IAM2ProjectTemplateInventoryView, error) {
	var resp view.IAM2ProjectTemplateInventoryView
	if err := cli.Get(ctx, "v1/iam2/projects/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2ProjectTemplate Pagination
func (cli *ZSClient) PageIAM2ProjectTemplate(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectTemplateInventoryView, int, error) {
	var iAM2ProjectTemplates []view.IAM2ProjectTemplateInventoryView
	total, err := cli.Page(ctx, "v1/iam2/projects/templates", params, &iAM2ProjectTemplates)
	return iAM2ProjectTemplates, total, err
}
// UpdateIAM2ProjectTemplate updates IAM2ProjectTemplate
func (cli *ZSClient) UpdateIAM2ProjectTemplate(ctx context.Context, uuid string, params param.UpdateIAM2ProjectTemplateParam) (*view.IAM2ProjectTemplateInventoryView, error) {
	resp := view.IAM2ProjectTemplateInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects/templates", uuid, "", map[string]interface{}{
		"updateIAM2ProjectTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIAM2ProjectTemplate deletes IAM2ProjectTemplate
func (cli *ZSClient) DeleteIAM2ProjectTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects/templates", uuid, string(deleteMode))
}

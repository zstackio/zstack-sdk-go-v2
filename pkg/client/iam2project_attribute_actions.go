// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2ProjectAttribute queries IAM2ProjectAttribute list
func (cli *ZSClient) QueryIAM2ProjectAttribute(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectAttributeInventoryView, error) {
	var resp []view.IAM2ProjectAttributeInventoryView
	return resp, cli.List(ctx, "v1/iam2/projects/attributes", params, &resp)
}

func (cli *ZSClient) GetIAM2ProjectAttribute(ctx context.Context, uuid string) (*view.IAM2ProjectAttributeInventoryView, error) {
	var resp view.IAM2ProjectAttributeInventoryView
	if err := cli.Get(ctx, "v1/iam2/projects/attributes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2ProjectAttribute Pagination
func (cli *ZSClient) PageIAM2ProjectAttribute(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectAttributeInventoryView, int, error) {
	var iAM2ProjectAttributes []view.IAM2ProjectAttributeInventoryView
	total, err := cli.Page(ctx, "v1/iam2/projects/attributes", params, &iAM2ProjectAttributes)
	return iAM2ProjectAttributes, total, err
}
// UpdateIAM2ProjectAttribute updates IAM2ProjectAttribute
func (cli *ZSClient) UpdateIAM2ProjectAttribute(ctx context.Context, uuid string, params param.UpdateIAM2ProjectAttributeParam) (*view.IAM2ProjectAttributeInventoryView, error) {
	resp := view.IAM2ProjectAttributeInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects/attributes", uuid, "", map[string]interface{}{
		"updateIAM2ProjectAttribute": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2ProjectAccountRef queries IAM2ProjectAccountRef list
func (cli *ZSClient) QueryIAM2ProjectAccountRef(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectAccountRefInventoryView, error) {
	var resp []view.IAM2ProjectAccountRefInventoryView
	return resp, cli.List(ctx, "v1/iam2/projects/account/refs", params, &resp)
}

func (cli *ZSClient) GetIAM2ProjectAccountRef(ctx context.Context, uuid string) (*view.IAM2ProjectAccountRefInventoryView, error) {
	var resp view.IAM2ProjectAccountRefInventoryView
	if err := cli.Get(ctx, "v1/iam2/projects/account/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2ProjectAccountRef Pagination
func (cli *ZSClient) PageIAM2ProjectAccountRef(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectAccountRefInventoryView, int, error) {
	var iAM2ProjectAccountRefs []view.IAM2ProjectAccountRefInventoryView
	total, err := cli.Page(ctx, "v1/iam2/projects/account/refs", params, &iAM2ProjectAccountRefs)
	return iAM2ProjectAccountRefs, total, err
}

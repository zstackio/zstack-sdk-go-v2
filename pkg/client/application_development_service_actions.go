// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryApplicationDevelopmentService queries ApplicationDevelopmentService list
func (cli *ZSClient) QueryApplicationDevelopmentService(ctx context.Context, params *param.QueryParam) ([]view.ApplicationDevelopmentServiceInventoryView, error) {
	var resp []view.ApplicationDevelopmentServiceInventoryView
	return resp, cli.List(ctx, "v1/ai/model-services/app/", params, &resp)
}

func (cli *ZSClient) GetApplicationDevelopmentService(ctx context.Context, uuid string) (*view.ApplicationDevelopmentServiceInventoryView, error) {
	var resp view.ApplicationDevelopmentServiceInventoryView
	if err := cli.Get(ctx, "v1/ai/model-services/app/", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageApplicationDevelopmentService Pagination
func (cli *ZSClient) PageApplicationDevelopmentService(ctx context.Context, params *param.QueryParam) ([]view.ApplicationDevelopmentServiceInventoryView, int, error) {
	var applicationDevelopmentServices []view.ApplicationDevelopmentServiceInventoryView
	total, err := cli.Page(ctx, "v1/ai/model-services/app/", params, &applicationDevelopmentServices)
	return applicationDevelopmentServices, total, err
}

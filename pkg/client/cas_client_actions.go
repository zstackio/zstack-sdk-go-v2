// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCasClient creates CasClient
func (cli *ZSClient) CreateCasClient(ctx context.Context, params param.CreateCasClientParam) (*view.CasClientInventoryView, error) {
	resp := view.CasClientInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/create/cas/client", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateCasClient updates CasClient
func (cli *ZSClient) UpdateCasClient(ctx context.Context, params param.UpdateCasClientParam) (*view.CasClientInventoryView, error) {
	resp := view.CasClientInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/update/cas/client", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

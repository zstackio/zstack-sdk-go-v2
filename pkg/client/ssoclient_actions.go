// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteSSOClient deletes SSOClient
func (cli *ZSClient) DeleteSSOClient(ctx context.Context, params param.DeleteSSOClientParam) (*view.SSOClientInventoryView, error) {
	resp := view.SSOClientInventoryView{}
	if err := cli.Post(ctx, "v1/delete/sso/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// GetSSOClient gets SSOClient by uuid
func (cli *ZSClient) GetSSOClient(ctx context.Context) (*view.SSOClientInventoryView, error) {
	var resp view.SSOClientInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/get/sso/client", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

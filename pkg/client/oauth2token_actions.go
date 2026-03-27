// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetOAuth2Token gets OAuth2Token by uuid
func (cli *ZSClient) GetOAuth2Token(ctx context.Context) (*view.OAuth2TokenInventoryView, error) {
	var resp view.GetOAuth2TokenView
	if err := cli.GetWithRespKey(ctx, "v1/get/oauth2/token", "", "inventory", nil, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

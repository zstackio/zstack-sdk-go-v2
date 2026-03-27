// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// RefreshSSOServerToken operates on SSOServerToken
func (cli *ZSClient) RefreshSSOServerToken(ctx context.Context, params param.RefreshSSOServerTokenParam) (*view.SSOServerTokenInventoryView, error) {
	resp := view.SSOServerTokenInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/sso/server/token/refresh", "", "", map[string]interface{}{
		"refreshSSOServerToken": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
